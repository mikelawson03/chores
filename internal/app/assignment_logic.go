package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mikelawson03/chores/internal/domain"
)

// Compare requesting user to user assignment and ensure they match (return error if not)
// Make and call helper to build Assignment add App.Assignments

func (a *App) validateTemplateID(ctx context.Context, templateID string) error {
	if strings.TrimSpace(templateID) == "" {
		return errors.New("missing required template ID")
	}

	_, err := a.GetChoreTemplateByID(ctx, templateID)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) validateAssignedUserID(ctx context.Context, assignedUserID string) error {
	if strings.TrimSpace(assignedUserID) == "" {
		return errors.New("missing required assigned user ID")
	}

	_, err := a.GetUserByID(ctx, assignedUserID)
	if err != nil {
		return err
	}

	return nil
}

func validateDueDate(dueDate *time.Time) error {
	if dueDate == nil {
		return errors.New("missing required schedule date")
	}

	if dueDate.Before(time.Now()) {
		return errors.New("chores may not be scheduled before current time")
	}
	return nil
}

func (a *App) createNewAssignment(templateID, assignedUserID, instructions string, dueDate *time.Time, scheduledFor *time.Time) domain.Assignment {
	id := uuid.NewString()
	now := time.Now()

	var assignmentScheduledFor *time.Time

	if scheduledFor != nil {
		t := *scheduledFor
		assignmentScheduledFor = &t
	}

	assignment := domain.Assignment{
		ID:             id,
		TemplateID:     templateID,
		AssignedUserID: assignedUserID,
		Instructions:   instructions,
		DueDate:        *dueDate,
		ScheduledFor:   assignmentScheduledFor,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	return assignment
}

func (a *App) CreateAssignmentForUser(ctx context.Context, requesterID, templateID, assignedUserID, instructions string, dueDate, scheduledFor *time.Time) (domain.Assignment, error) {
	// validate that templateID is provided and exists
	err := a.validateTemplateID(ctx, templateID)
	if err != nil {
		return domain.Assignment{}, err
	}

	// validate that assignedUserID is provided and exists
	err = a.validateAssignedUserID(ctx, assignedUserID)
	if err != nil {
		return domain.Assignment{}, err
	}

	// validate schedule date exists and is in the future
	err = validateDueDate(dueDate)
	if err != nil {
		return domain.Assignment{}, err
	}

	// check authorization (may only assign to self)
	if requesterID != assignedUserID {
		return domain.Assignment{}, errors.New("may not assign chores to other users")
	}

	// check that chore template exists
	_, err = a.GetChoreTemplateByID(ctx, templateID)
	if err != nil {
		return domain.Assignment{}, err
	}

	assignment := a.createNewAssignment(templateID, assignedUserID, instructions, dueDate, scheduledFor)

	err = a.Store.AddAssignment(ctx, assignment)
	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
}

func (a *App) GetAllAssignments(ctx context.Context) ([]domain.Assignment, error) {
	assignments, err := a.Store.GetAllAssignments(ctx)
	if err != nil {
		return []domain.Assignment{}, err
	}

	return assignments, nil
}

func (a *App) GetAssignmentByID(ctx context.Context, id string) (domain.Assignment, error) {
	assignment, err := a.Store.GetAssignment(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Assignment{}, fmt.Errorf("assignment with ID %s not found", id)
	}

	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
}

func (a *App) EditAssignment(ctx context.Context, requesterID, assignmentID, assignedUserID, notes string, canceled, completed bool, scheduledFor *time.Time) (domain.Assignment, error) {
	// check assignment exists
	existing, err := a.GetAssignmentByID(ctx, assignmentID)
	if err != nil {
		return domain.Assignment{}, err
	}

	// // check user owns assignment
	// if existing.AssignedUserID != requesterID {
	// 	return domain.Assignment{}, errors.New("may not edit chores belonging to other users")
	// }

	// check that userID is provided and valid
	// err = a.validateAssignedUserID(ctx, assignedUserID)
	// if err != nil {
	// 	return domain.Assignment{}, err
	// }

	// // check assignment is to self
	// if requesterID != assignedUserID {
	// 	return domain.Assignment{}, errors.New("may not assign chores to other users")
	// }

	var completedAt *time.Time
	var canceledAt *time.Time

	if !existing.Completed && completed {
		t := time.Now()
		completedAt = &t
	} else if existing.Completed && !completed {
		completedAt = nil
	}

	if !existing.Canceled && canceled {
		t := time.Now()
		canceledAt = &t
	} else if existing.Canceled && !canceled {
		canceledAt = nil
	}

	assignment := domain.Assignment{
		ID:             assignmentID,
		TemplateID:     existing.TemplateID,
		TemplateName:   existing.TemplateName,
		AssignedUserID: assignedUserID,
		Cadence:        existing.Cadence,
		Duration:       existing.Duration,
		Instructions:   existing.Instructions,
		Notes:          notes,
		DueDate:        existing.DueDate,
		ScheduledFor:   scheduledFor,
		Completed:      completed,
		Canceled:       canceled,
		CreatedAt:      existing.CreatedAt,
		UpdatedAt:      time.Now(),
		CompletedAt:    completedAt,
		CanceledAt:     canceledAt,
	}

	err = a.Store.EditAssignment(ctx, assignment)
	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
}

func (a *App) GetAssignmentsByTemplateID(ctx context.Context, id string) ([]domain.Assignment, error) {
	_, err := a.GetChoreTemplateByID(ctx, id)
	if err != nil {
		return []domain.Assignment{}, err
	}

	assignments, err := a.Store.GetAssignmentsByTemplateID(ctx, id)
	if err != nil {
		return []domain.Assignment{}, err
	}

	return assignments, nil
}

func (a *App) GetAssignmentsByUserID(ctx context.Context, id string) ([]domain.Assignment, error) {
	_, err := a.GetUserByID(ctx, id)
	if err != nil {
		return []domain.Assignment{}, err
	}

	assignments, err := a.Store.GetAssignmentsByUserID(ctx, id)
	if err != nil {
		return []domain.Assignment{}, err
	}

	return assignments, nil
}
