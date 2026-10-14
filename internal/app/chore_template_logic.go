package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	domain "github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
)

func newChoreTemplateEvent(eventType domain.ChoreTemplateEventType, occurredAt time.Time, actorType domain.ActorType, choreTemplate domain.ChoreTemplate, householdID, userID string) domain.Event[domain.ChoreTemplate] {
	e := domain.Event[domain.ChoreTemplate]{
		Type:        string(eventType),
		OccurredAt:  occurredAt,
		HouseholdID: householdID,
		Actor: domain.Actor{
			Type: actorType,
			ID:   userID,
		},
		Payload: choreTemplate,
	}

	return e
}

func (a *App) choreNameExists(ctx context.Context, name, id string) error {
	ct, err := a.Store.GetTemplateByName(ctx, name)

	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}

	if err != nil {
		return err
	}

	if id != "" && ct.ID == id {
		return nil
	}

	return fmt.Errorf("%w: chore template name already exists", domain.ErrConflict)
}

func (a *App) validateChoreTemplateRequest(ctx context.Context, name, cadence, assignee, id, householdId string, duration *int) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("%w: name required", domain.ErrInvalidRequest)
	}

	if len(name) > 50 {
		return fmt.Errorf("%w: name must be 50 characters or fewer", domain.ErrInvalidRequest)
	}

	if strings.TrimSpace(cadence) == "" {
		return fmt.Errorf("%w: cadence required", domain.ErrInvalidRequest)
	}

	dc := domain.Cadence(cadence)
	if !dc.IsValid() {
		return domain.ErrInvalidCadence
	}

	if duration == nil || *duration <= 0 {
		return fmt.Errorf("%w: duration must be greater than zero", domain.ErrInvalidRequest)
	}

	if assignee != "" {
		err := a.ValidateChoreTemplateAssignee(ctx, assignee, householdId)
		if err != nil {
			return err
		}
	}

	err := a.choreNameExists(ctx, name, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) ValidateChoreTemplateAssignee(ctx context.Context, assigneeID, householdId string) error {
	_, err := a.Store.GetHouseholdUserByID(ctx, assigneeID, householdId)
	if errors.Is(err, domain.ErrNotFound) {
		return fmt.Errorf("%w: assigned user not found", domain.ErrInvalidRequest)
	}
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetChoreTemplates(ctx context.Context) ([]domain.ChoreTemplate, error) {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return []domain.ChoreTemplate{}, err
	}

	tmps, err := a.Store.GetChoreTemplates(ctx)
	if err != nil {
		return []domain.ChoreTemplate{}, err
	}
	return tmps, nil
}

func (a *App) GetChoreTemplateByID(ctx context.Context, id string) (domain.ChoreTemplate, error) {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	tmp, err := a.Store.GetTemplateByID(ctx, id)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	return tmp, nil
}

func (a *App) CreateChoreTemplate(ctx context.Context, name, cadence, assignee, instructions string, duration *int) (domain.ChoreTemplate, error) {

	hhUser, err := CheckAdmin(ctx)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	err = a.validateChoreTemplateRequest(ctx, name, cadence, assignee, "", hhUser.HouseholdID, duration)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	id := uuid.NewString()
	now := time.Now()

	chore := domain.ChoreTemplate{
		ID:           id,
		Name:         name,
		Cadence:      domain.Cadence(cadence),
		Assignee:     assignee,
		Instructions: instructions,
		IsActive:     true,
		Duration:     *duration,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err = a.Store.AddChoreTemplate(ctx, chore)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	event := newChoreTemplateEvent(domain.ChoreTemplateCreated,
		now,
		domain.ActorTypeUser,
		chore,
		domain.DefaultHouseholdID,
		hhUser.User.ID,
	)

	a.Bus.Publish(event)

	return chore, nil
}

func (a *App) EditChoreTemplate(ctx context.Context, id, name, cadence, assignee, instructions string, duration *int) (domain.ChoreTemplate, error) {
	hhUser, err := CheckAdmin(ctx)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	tmp, err := a.Store.GetTemplateByID(ctx, id)
	if err != nil {
		return domain.ChoreTemplate{}, fmt.Errorf("%w: chore template", domain.ErrNotFound)
	}

	err = a.validateChoreTemplateRequest(ctx, name, cadence, assignee, id, hhUser.HouseholdID, duration)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	now := time.Now()

	updatedTmp := domain.ChoreTemplate{
		ID:           id,
		Name:         name,
		Cadence:      domain.Cadence(cadence),
		Assignee:     assignee,
		Instructions: instructions,
		Duration:     *duration,
		IsActive:     tmp.IsActive,
		CreatedAt:    tmp.CreatedAt,
		UpdatedAt:    now,
	}

	err = a.Store.EditChoreTemplate(ctx, updatedTmp)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	event := newChoreTemplateEvent(domain.ChoreTemplateEdited,
		now,
		domain.ActorTypeUser,
		updatedTmp,
		hhUser.HouseholdID,
		hhUser.User.ID,
	)

	a.Bus.Publish(event)

	return updatedTmp, nil
}

func (a *App) DeactivateChoreTemplate(ctx context.Context, id string) error {
	reqUser, err := CheckAdmin(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	var deactivatedTemplate domain.ChoreTemplate
	var canceledAssignments []domain.Assignment

	err = a.Store.WithTx(ctx, func(txStore *store.Store) error {
		deactivatedTemplate, err = txStore.DeactivateChoreTemplate(ctx, id, now)
		if errors.Is(err, domain.ErrNotFound) {
			return fmt.Errorf("%w: chore template", err)
		}

		if err != nil {
			return err
		}

		canceledAssignments, err = txStore.CancelOutstandingAssignmentsByTemplateID(ctx, &now, now, id)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	event := newChoreTemplateEvent(
		domain.ChoreTemplateDeactivated,
		now,
		domain.ActorTypeUser,
		deactivatedTemplate,
		reqUser.HouseholdID,
		reqUser.User.ID,
	)

	a.Bus.Publish(event)

	for _, assignment := range canceledAssignments {
		event := newAssignmentEvent(
			domain.AssignmentEdited,
			now,
			domain.ActorTypeUser,
			assignment,
			reqUser.HouseholdID,
			reqUser.User.ID,
		)

		a.Bus.Publish(event)
	}

	return nil
}
