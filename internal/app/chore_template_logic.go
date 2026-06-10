package app

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	domain "github.com/mikelawson03/chores/internal/domain"
)

func (a *App) choreExists(id string) bool {
	if _, exists := a.Templates[id]; exists {
		return true
	}
	return false
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

func (a *App) CreateChoreTemplate(name string, cadence string, shared *bool, assignee string, duration *int) (domain.ChoreTemplate, error) {
	if a.choreExists(name) {
		return domain.ChoreTemplate{}, fmt.Errorf("Chore with name %s already exists", name)
	}

	err := a.validateChoreTemplateRequest(name, cadence, shared, duration)
	if err != nil {
		return domain.ChoreTemplate{}, err
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

	err = a.Store.CreateTemplate(context.Background(), chore)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	return chore, nil
}

func (a *App) GetChoreTemplates() []domain.ChoreTemplate {
	var chores []domain.ChoreTemplate
	for _, chore := range a.Templates {
		chores = append(chores, chore)
	}
	return chores
}

func (a *App) GetChoreByID(id string) (domain.ChoreTemplate, error) {
	if a.choreExists(id) {
		return a.Templates[id], nil
	}

	return domain.ChoreTemplate{}, fmt.Errorf("Chore `%s` not found.", id)
}

func (a *App) EditChoreTemplate(id string, name string, cadence string, shared *bool, assignee string, duration *int) (domain.ChoreTemplate, error) {
	if a.choreExists(id) {
		err := a.validateChoreTemplateRequest(name, cadence, shared, duration)
		if err != nil {
			return domain.ChoreTemplate{}, err
		}

		chore := domain.ChoreTemplate{
			ID:        id,
			Name:      name,
			Cadence:   cadence,
			Shared:    *shared,
			Assignee:  assignee,
			Duration:  *duration,
			CreatedAt: a.Templates[id].CreatedAt,
			UpdatedAt: time.Now(),
		}

		a.Templates[id] = chore

		return chore, nil
	}

	return domain.ChoreTemplate{}, fmt.Errorf("Chore `%s` not found.", name)
}

func (a *App) DeleteChoreTemplate(id string) error {
	if !a.choreExists(id) {
		return fmt.Errorf("Chore ID %s not found", id)
	}
	delete(a.Templates, id)
	return nil
}

func (a *App) GetAssignmentsByTemplateID(id string) ([]domain.Assignment, error) {
	assignments := make([]domain.Assignment, 0)

	if !a.choreExists(id) {
		return assignments, fmt.Errorf("Chore template %s not found", id)
	}

	for _, assignment := range a.Assignments {
		if assignment.TemplateID == id {
			assignments = append(assignments, assignment)
		}
	}

	return assignments, nil
}
