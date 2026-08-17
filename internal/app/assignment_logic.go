package app

import (
	"context"
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
		return fmt.Errorf("%w: template ID required", domain.ErrInvalidRequest)
	}

	_, err := a.Store.GetTemplateByID(ctx, templateID)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) validateAssignedUserID(ctx context.Context, user domain.User, assignedUserID string) error {

	_, err := a.Store.GetUserByID(ctx, assignedUserID)
	if errors.Is(err, domain.ErrNotFound) {
		return fmt.Errorf("%w: assigned user not found", domain.ErrInvalidRequest)
	}
	if err != nil {
		return err
	}

	if !auth.CanAssignToUser(user, assignedUserID) {
		return fmt.Errorf("%w: may only assign chores to self", domain.ErrForbidden)
	}

	return nil
}

func validateDueDate(dueDate *time.Time) error {
	if dueDate == nil {
		return fmt.Errorf("%w: due date required", domain.ErrInvalidRequest)
	}

	if dueDate.Before(time.Now()) {
		return fmt.Errorf("%w: due date cannot be in the past", domain.ErrInvalidRequest)
	}
	return nil
}

func (a *App) validateCreateAssignment(ctx context.Context, user domain.User, req CreateAsssignmentRequest) error {
	err := a.validateTemplateID(ctx, req.TemplateID)
	if err != nil {
		return err
	}

	err = a.validateAssignedUserID(ctx, user, req.AssignedUserID)
	if err != nil {
		return err
	}

	err = validateDueDate(req.DueDate)
	if err != nil {
		return err
	}

	if req.ScheduledFor != nil && req.DueDate.Before(*req.ScheduledFor) {
		return fmt.Errorf("%w: cannot schedule assignment after due date", domain.ErrInvalidRequest)
	}

	return nil
}

func (a *App) validateEditAssignmentRequest(ctx context.Context, user domain.User, req EditAssignmentRequest, existing domain.Assignment) error {
	if !auth.CanEditAssignment(user, existing) {
		return domain.ErrForbidden
	}

	err := a.validateAssignedUserID(ctx, user, req.AssignedUserID)
	if err != nil {
		return err
	}

	if !auth.CanAssignToUser(user, req.AssignedUserID) {
		return domain.ErrForbidden
	}

	if req.ScheduledFor != nil && !req.Completed && existing.DueDate.Before(*req.ScheduledFor) {
		return fmt.Errorf("%w: cannot schedule assignment after due date", domain.ErrInvalidRequest)
	}

	return nil
}

func (a *App) CreateAssignmentFromTemplate(ctx context.Context, req CreateAsssignmentRequest) (domain.Assignment, error) {
	user, err := auth.AuthenticatedUser(ctx)
	if err != nil {
		return domain.Assignment{}, err
	}

	if err := a.validateCreateAssignment(ctx, user, req); err != nil {
		return domain.Assignment{}, err
	}

	id := uuid.NewString()
	now := time.Now()

	err = a.Store.AddAssignment(ctx, store.CreateAssignmentParams{
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

	assignment, err := a.Store.GetAssignment(ctx, id)
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
	user, err := auth.AuthenticatedUser(ctx)
	if err != nil {
		return domain.Assignment{}, err
	}

	assignment, err := a.Store.GetAssignment(ctx, id)
	if err != nil {
		return domain.Assignment{}, err
	}

	if !auth.CanGetAssignment(user, assignment) {
		return domain.Assignment{}, domain.ErrForbidden
	}

	return assignment, nil
}

func (a *App) EditAssignment(ctx context.Context, editRequest EditAssignmentRequest) (domain.Assignment, error) {
	user, err := auth.AuthenticatedUser(ctx)
	if err != nil {
		return domain.Assignment{}, err
	}

	existing, err := a.Store.GetAssignment(ctx, editRequest.ID)
	if err != nil {
		return domain.Assignment{}, err
	}

	err = a.validateEditAssignmentRequest(ctx, user, editRequest, existing)
	if err != nil {
		return domain.Assignment{}, err
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

	_, err = a.Store.GetTemplateByID(ctx, id)
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
	user, err := auth.AuthenticatedUser(ctx)
	if err != nil {
		return []domain.Assignment{}, err
	}

	if !auth.CanGetUserAssignments(user, id) {
		return []domain.Assignment{}, fmt.Errorf("%w: may only retrieve own assignments", domain.ErrForbidden)
	}

	_, err = a.Store.GetUserByID(ctx, id)
	if err != nil {
		return []domain.Assignment{}, err
	}

	assignments, err := a.Store.GetAssignmentsByUserID(ctx, id)
	if err != nil {
		return []domain.Assignment{}, err
	}

	return assignments, nil
}
