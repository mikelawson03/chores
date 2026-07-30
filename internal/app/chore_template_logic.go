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

func (a *App) validateChoreTemplateRequest(name string, cadence string, duration *int) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(cadence) == "" {
		return errors.New("cadence is required")
	}
	if duration == nil {
		return errors.New("duration is required")
	}
	return nil
}

func (a *App) ValidateChoreTemplateAssignee(ctx context.Context, newAssignee string) error {
	_, err := a.GetUserByID(ctx, newAssignee)
	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("assigned user does not exist")
	}

	return nil
}

func (a *App) CreateChoreTemplate(ctx context.Context, name, cadence, assignee, instructions string, duration *int) (domain.ChoreTemplate, error) {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	exists, err := a.choreNameExists(ctx, name)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	if exists {
		return domain.ChoreTemplate{}, fmt.Errorf("chore with name %s already exists", name)
	}

	err = a.validateChoreTemplateRequest(name, cadence, duration)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	if assignee != "" {
		err = a.ValidateChoreTemplateAssignee(ctx, assignee)
		if err != nil {
			return domain.ChoreTemplate{}, err
		}
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

	err = a.Store.AddChoreTemplate(context.Background(), chore)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	return chore, nil
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

	if errors.Is(err, sql.ErrNoRows) {
		err = fmt.Errorf("%w: template", domain.ErrNotFound)
	}

	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	return tmp, nil
}

func (a *App) EditChoreTemplate(ctx context.Context, id, name, cadence, assignee, instructions string, duration *int) (domain.ChoreTemplate, error) {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	tmp, err := a.GetChoreTemplateByID(ctx, id)

	if err != nil {
		return tmp, err
	}

	if assignee != "" {
		err = a.ValidateChoreTemplateAssignee(ctx, assignee)
		if err != nil {
			return domain.ChoreTemplate{}, err
		}
	}

	err = a.validateChoreTemplateRequest(name, cadence, duration)
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
