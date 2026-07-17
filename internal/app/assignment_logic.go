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
		return errors.New("Missing required template ID")
	}

	_, err := a.GetChoreTemplateByID(ctx, templateID)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) validateAssignedUserID(ctx context.Context, assignedUserID string) error {
	if strings.TrimSpace(assignedUserID) == "" {
		return errors.New("Missing required assigned user ID")
	}

	_, err := a.GetUserByID(ctx, assignedUserID)
	if err != nil {
		return err
	}

	return nil
}

func validateDueDate(dueDate *time.Time) error {
	if dueDate == nil {
		return errors.New("Missing required schedule date")
	}

	if dueDate.Before(time.Now()) {
		return errors.New("Chores may not be scheduled before current time")
	}
	return nil
}

func (a *App) createNewAssignment(templateID string, assignedUserID string, dueDate *time.Time) domain.Assignment {
	id := uuid.NewString()
	now := time.Now()
	assignment := domain.Assignment{
		ID:             id,
		TemplateID:     templateID,
		AssignedUserID: assignedUserID,
		DueDate:        *dueDate,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	return assignment
}

func (a *App) CreateAssignmentForUser(ctx context.Context, requesterID, templateID, assignedUserID string, dueDate *time.Time) (domain.Assignment, error) {
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

	assignment := a.createNewAssignment(templateID, assignedUserID, dueDate)

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
	assignment, err := a.Store.GetAssignmentByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Assignment{}, fmt.Errorf("Assignment with ID %s not found.", id)
	}

	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
}

func (a *App) EditAssignment(ctx context.Context, requesterID, assignmentID, assignedUserID string) (domain.Assignment, error) {
	// check assignment exists
	existing, err := a.GetAssignmentByID(ctx, assignmentID)
	if err != nil {
		return domain.Assignment{}, err
	}

	// check user owns assignment
	if existing.AssignedUserID != requesterID {
		return domain.Assignment{}, errors.New("may not edit chores belonging to other users")
	}

	// check that userID is provided and valid
	err = a.validateAssignedUserID(ctx, assignedUserID)
	if err != nil {
		return domain.Assignment{}, err
	}

	// check assignment is to self
	if requesterID != assignedUserID {
		return domain.Assignment{}, errors.New("may not assign chores to other users")
	}

	assignment := domain.Assignment{
		ID:             assignmentID,
		TemplateID:     existing.TemplateID,
		AssignedUserID: assignedUserID,
		Completed:      existing.Completed,
		Canceled:       existing.Canceled,
		CreatedAt:      existing.CreatedAt,
		UpdatedAt:      time.Now(),
		CompletedAt:    existing.CompletedAt,
		CanceledAt:     existing.CanceledAt,
	}

	err = a.Store.EditAssignment(ctx, assignment)
	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
}

func (a *App) CancelAssignment(ctx context.Context, assignmentID, requesterID string) (domain.Assignment, error) {
	existing, err := a.GetAssignmentByID(ctx, assignmentID)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Assignment{}, errors.New("Assignment not found")
	}

	if err != nil {
		return domain.Assignment{}, err
	}

	if existing.AssignedUserID != requesterID {
		return domain.Assignment{}, errors.New("May not cancel other user's assignments")
	}

	now := time.Now()

	assignment := domain.Assignment{
		ID:             assignmentID,
		TemplateID:     existing.TemplateID,
		AssignedUserID: existing.AssignedUserID,
		DueDate:        existing.DueDate,
		Completed:      existing.Completed,
		Canceled:       true,
		CreatedAt:      existing.CreatedAt,
		UpdatedAt:      now,
		CompletedAt:    existing.CompletedAt,
		CanceledAt:     now,
	}

	err = a.Store.CancelAssignment(ctx, assignment)
	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
}

func (a *App) CompleteAssignment(ctx context.Context, assignmentID, requesterID string) (domain.Assignment, error) {
	existing, err := a.GetAssignmentByID(ctx, assignmentID)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Assignment{}, errors.New("Assignment not found")
	}

	if err != nil {
		return domain.Assignment{}, err
	}

	if existing.AssignedUserID != requesterID {
		return domain.Assignment{}, errors.New("May not complete other user's assignments")
	}
	now := time.Now()

	assignment := domain.Assignment{
		ID:             assignmentID,
		TemplateID:     existing.TemplateID,
		AssignedUserID: existing.AssignedUserID,
		DueDate:        existing.DueDate,
		Completed:      true,
		Canceled:       existing.Canceled,
		CreatedAt:      existing.CreatedAt,
		UpdatedAt:      now,
		CompletedAt:    now,
		CanceledAt:     existing.CanceledAt,
	}

	err = a.Store.CancelAssignment(ctx, assignment)
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
