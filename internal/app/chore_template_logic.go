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

func (a *App) choreNameExists(ctx context.Context, name string) (bool, error) {
	_, err := a.Store.GetTemplateByName(ctx, name)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *App) validateChoreTemplateRequest(name string, cadence string, shared *bool, duration *int) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("Name is required")
	}
	if strings.TrimSpace(cadence) == "" {
		return errors.New("Cadence is required")
	}
	if shared == nil {
		return errors.New("Sharing value required")
	}
	if duration == nil {
		return errors.New("Duration is required")
	}
	return nil
}

func (a *App) ValidateChoreTemplateAssignee(ctx context.Context, requesterID, newAssignee string) error {
	_, err := a.GetUserByID(ctx, newAssignee)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("Assignment invalid - assigned user does not exist")
	}

	if newAssignee != "" && requesterID != newAssignee {
		return errors.New("May not assign templates to other users")
	}

	return nil
}

func (a *App) CreateChoreTemplate(ctx context.Context, requesterID, name, cadence, assignee string, shared *bool, duration *int) (domain.ChoreTemplate, error) {
	exists, err := a.choreNameExists(ctx, name)

	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	if exists {
		return domain.ChoreTemplate{}, fmt.Errorf("Chore with name %s already exists", name)
	}

	err = a.validateChoreTemplateRequest(name, cadence, shared, duration)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	if assignee != "" {
		err = a.ValidateChoreTemplateAssignee(ctx, requesterID, assignee)
		if err != nil {
			return domain.ChoreTemplate{}, err
		}
	}

	id := uuid.NewString()

	chore := domain.ChoreTemplate{
		ID:        id,
		Name:      name,
		Cadence:   cadence,
		Shared:    *shared,
		Assignee:  assignee,
		Duration:  *duration,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = a.Store.AddChoreTemplate(context.Background(), chore)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	return chore, nil
}

func (a *App) GetChoreTemplates(ctx context.Context) ([]domain.ChoreTemplate, error) {
	tmps, err := a.Store.GetChoreTemplates(ctx)
	if err != nil {
		return []domain.ChoreTemplate{}, err
	}
	return tmps, nil
}

func (a *App) GetChoreTemplateByID(ctx context.Context, id string) (domain.ChoreTemplate, error) {
	tmp, err := a.Store.GetTemplateByID(ctx, id)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.ChoreTemplate{}, fmt.Errorf("Chore `%s` not found.", id)
	}

	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	return tmp, nil
}

func (a *App) EditChoreTemplate(ctx context.Context, requesterID, id, name, cadence, assignee string, shared *bool, duration *int) (domain.ChoreTemplate, error) {
	tmp, err := a.GetChoreTemplateByID(ctx, id)

	if err != nil {
		return tmp, err
	}

	if tmp.Assignee != "" && tmp.Assignee != requesterID {
		return domain.ChoreTemplate{}, errors.New("May not edit chores assigned to other users")
	}

	if assignee != "" {
		err = a.ValidateChoreTemplateAssignee(ctx, requesterID, assignee)
		if err != nil {
			return domain.ChoreTemplate{}, err
		}
	}

	err = a.validateChoreTemplateRequest(name, cadence, shared, duration)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	updatedTmp := domain.ChoreTemplate{
		ID:        id,
		Name:      name,
		Cadence:   cadence,
		Shared:    *shared,
		Assignee:  assignee,
		Duration:  *duration,
		CreatedAt: tmp.CreatedAt,
		UpdatedAt: time.Now(),
	}

	err = a.Store.EditChoreTemplate(ctx, updatedTmp)

	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	return updatedTmp, nil
}

func (a *App) DeleteChoreTemplate(ctx context.Context, id, requesterID string) error {
	tmp, err := a.GetChoreTemplateByID(ctx, id)
	if err != nil {
		return err
	}

	if tmp.Assignee != "" && requesterID != tmp.Assignee {
		return errors.New("May not delete other users' templates")
	}

	err = a.Store.DeleteChoreTemplate(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) GetAssignmentsByTemplateID(ctx context.Context, id string) ([]domain.Assignment, error) {
	assignments := make([]domain.Assignment, 0)

	_, err := a.GetChoreTemplateByID(ctx, id)
	if err != nil {
		return []domain.Assignment{}, err
	}

	for _, assignment := range a.Assignments {
		if assignment.TemplateID == id {
			assignments = append(assignments, assignment)
		}
	}

	return assignments, nil
}
