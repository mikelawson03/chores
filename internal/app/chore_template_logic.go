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
)

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

func (a *App) validateChoreTemplateRequest(ctx context.Context, name, cadence, assignee, id string, duration *int) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("%w: name is required", domain.ErrInvalidRequest)
	}
	if strings.TrimSpace(cadence) == "" {
		return fmt.Errorf("%w: cadence is required", domain.ErrInvalidRequest)
	}

	dc := domain.Cadence(cadence)
	if !dc.IsValid() {
		return domain.ErrInvalidCadence
	}

	if duration == nil {
		return fmt.Errorf("%w: duration is required", domain.ErrInvalidRequest)
	}
	if assignee != "" {
		err := a.ValidateChoreTemplateAssignee(ctx, assignee)
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

func (a *App) ValidateChoreTemplateAssignee(ctx context.Context, assigneeID string) error {
	_, err := a.Store.GetUserByID(ctx, assigneeID)
	if errors.Is(err, domain.ErrNotFound) {
		return fmt.Errorf("%w: assigned user does not exist", domain.ErrInvalidRequest)
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
	_, err := CheckAdmin(ctx)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	err = a.validateChoreTemplateRequest(ctx, name, cadence, assignee, "", duration)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	id := uuid.NewString()

	chore := domain.ChoreTemplate{
		ID:           id,
		Name:         name,
		Cadence:      domain.Cadence(cadence),
		Assignee:     assignee,
		Instructions: instructions,
		Duration:     *duration,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = a.Store.AddChoreTemplate(ctx, chore)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	return chore, nil
}

func (a *App) EditChoreTemplate(ctx context.Context, id, name, cadence, assignee, instructions string, duration *int) (domain.ChoreTemplate, error) {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	tmp, err := a.Store.GetTemplateByID(ctx, id)
	if err != nil {
		return domain.ChoreTemplate{}, fmt.Errorf("%w: chore template", domain.ErrNotFound)
	}

	err = a.validateChoreTemplateRequest(ctx, name, cadence, assignee, id, duration)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	updatedTmp := domain.ChoreTemplate{
		ID:           id,
		Name:         name,
		Cadence:      domain.Cadence(cadence),
		Assignee:     assignee,
		Instructions: instructions,
		Duration:     *duration,
		CreatedAt:    tmp.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	err = a.Store.EditChoreTemplate(ctx, updatedTmp)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	return updatedTmp, nil
}

func (a *App) DeleteChoreTemplate(ctx context.Context, id string) error {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return err
	}

	err = a.Store.DeleteChoreTemplate(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w: chore template", domain.ErrNotFound)
	}

	if err != nil {
		return err
	}
	return nil
}
