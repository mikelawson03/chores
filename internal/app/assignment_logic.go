package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mikelawson03/chores/internal/auth"
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
)

type CreateAsssignmentRequest struct {
	TemplateID     string
	AssignedUserID string
	Instructions   string
	DueDate        *time.Time
	ScheduledFor   *time.Time
}

type EditAssignmentRequest struct {
	ID             string
	AssignedUserID string
	ScheduledFor   *time.Time
	Notes          string
	Completed      bool
	Canceled       bool
}

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

	_, err := a.GetUserByID(ctx, assignedUserID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w: assigned user", ErrNotFound)
	}
	if err != nil {
		return err
	}

	return nil
}

func (a *App) validateCreateAssignment(ctx context.Context, req CreateAsssignmentRequest) error {
	err := a.validateTemplateID(ctx, req.TemplateID)
	if err != nil {
		return err
	}

	err = a.validateAssignedUserID(ctx, req.AssignedUserID)
	if err != nil {
		return err
	}

	err = validateDueDate(req.DueDate)
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

func (a *App) CreateAssignmentFromTemplate(ctx context.Context, req CreateAsssignmentRequest) (domain.Assignment, error) {
	user, err := AuthenticatedUser(ctx)
	if err != nil {
		return domain.Assignment{}, err
	}

	if err := a.validateCreateAssignment(ctx, req); err != nil {
		return domain.Assignment{}, err
	}

	if !auth.CanAssignToUser(user, req.AssignedUserID) {
		return domain.Assignment{}, fmt.Errorf("%w: may only assign chores to self", ErrForbidden)
	}

	id := uuid.NewString()
	now := time.Now()

	assignment, err := a.Store.AddAssignment(ctx, store.CreateAssignmentParams{
		ID:             id,
		TemplateID:     req.TemplateID,
		AssignedUserID: req.AssignedUserID,
		Instructions:   req.Instructions,
		DueDate:        *req.DueDate,
		ScheduledFor:   req.ScheduledFor,
		CreatedAt:      now,
		UpdatedAt:      now,
	})
	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
}

func (a *App) GetAllAssignments(ctx context.Context) ([]domain.Assignment, error) {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return []domain.Assignment{}, err
	}

	assignments, err := a.Store.GetAllAssignments(ctx)
	if err != nil {
		return []domain.Assignment{}, err
	}

	return assignments, nil
}

func (a *App) GetAssignmentByID(ctx context.Context, id string) (domain.Assignment, error) {
	user, err := AuthenticatedUser(ctx)
	if err != nil {
		return domain.Assignment{}, err
	}

	assignment, err := a.Store.GetAssignment(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Assignment{}, fmt.Errorf("assignment with ID %s not found", id)
	}

	if err != nil {
		return domain.Assignment{}, err
	}

	if !auth.CanGetAssignment(user, assignment) {
		return domain.Assignment{}, ErrForbidden
	}

	return assignment, nil
}

func (a *App) EditAssignment(ctx context.Context, editRequest EditAssignmentRequest) (domain.Assignment, error) {

	user, err := AuthenticatedUser(ctx)
	if err != nil {
		return domain.Assignment{}, err
	}

	// check assignment exists
	existing, err := a.Store.GetAssignment(ctx, editRequest.ID)
	if err != nil {
		return domain.Assignment{}, err
	}

	// check that userID is provided and valid
	err = a.validateAssignedUserID(ctx, editRequest.AssignedUserID)
	if err != nil {
		return domain.Assignment{}, err
	}

	// check authorization
	if !auth.CanEditAssignment(user, existing) {
		return domain.Assignment{}, ErrForbidden
	}

	if !auth.CanAssignToUser(user, editRequest.AssignedUserID) {
		return domain.Assignment{}, ErrForbidden
	}

	var completedAt *time.Time
	var canceledAt *time.Time

	if existing.Completed != editRequest.Completed {
		if editRequest.Completed {
			t := time.Now()
			completedAt = &t
		} else {
			completedAt = nil
		}
	} else {
		completedAt = existing.CompletedAt
	}

	if existing.Canceled != editRequest.Canceled {
		if editRequest.Canceled {
			t := time.Now()
			canceledAt = &t
		} else {
			canceledAt = nil
		}
	} else {
		canceledAt = existing.CanceledAt
	}

	assignment, err := a.Store.EditAssignment(ctx, store.EditAssignmentParams{
		ID:             editRequest.ID,
		AssignedUserID: editRequest.AssignedUserID,
		Notes:          editRequest.Notes,
		ScheduledFor:   editRequest.ScheduledFor,
		Completed:      editRequest.Completed,
		Canceled:       editRequest.Canceled,
		CompletedAt:    completedAt,
		CanceledAt:     canceledAt,
		UpdatedAt:      time.Now(),
	})
	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
}

func (a *App) GetAssignmentsByTemplateID(ctx context.Context, id string) ([]domain.Assignment, error) {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return []domain.Assignment{}, err
	}

	_, err = a.GetChoreTemplateByID(ctx, id)
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
	user, err := AuthenticatedUser(ctx)
	if err != nil {
		return []domain.Assignment{}, err
	}

	if !auth.CanGetUserAssignments(user, id) {
		return []domain.Assignment{}, fmt.Errorf("%w: may only get own assignments", ErrForbidden)
	}

	assignments, err := a.Store.GetAssignmentsByUserID(ctx, id)
	if err != nil {
		return []domain.Assignment{}, err
	}

	return assignments, nil
}
