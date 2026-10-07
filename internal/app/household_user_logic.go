package app

import (
	"context"
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/mikelawson03/chores/internal/auth"
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
)

type HouseholdUserRequest struct {
	Role        string
	DisplayName string
	ColorOption int
	IsActive    bool
	UserID      string
	HouseholdID string
}

func newHouseholdUserEvent(eventType domain.HouseholdUserEventType,
	occurredAt time.Time,
	actorType domain.ActorType,
	hhUser domain.HouseholdUser,
	householdID,
	userID string) domain.Event[domain.HouseholdUser] {
	return domain.Event[domain.HouseholdUser]{
		Type:        string(eventType),
		OccurredAt:  occurredAt,
		HouseholdID: householdID,
		Actor: domain.Actor{
			Type: actorType,
			ID:   userID,
		},
		Payload: hhUser,
	}
}

func (a *App) HouseholdUserExists(ctx context.Context, userID, householdID string) (bool, error) {
	_, err := a.Store.GetHouseholdUserByID(ctx, userID, householdID)
	if errors.Is(err, domain.ErrNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *App) validateHouseholdUserFields(ctx context.Context, req HouseholdUserRequest) error {
	role := domain.Role(req.Role)
	if !role.IsValid() {
		return domain.ErrInvalidRole
	}

	if req.ColorOption < 0 || req.ColorOption >= domain.MaxHouseholdMembers {
		return domain.ErrInvalidColorOption
	}

	if utf8.RuneCountInString(req.DisplayName) > 16 {
		return domain.ErrInvalidDisplayName
	}

	exists, err := a.Store.HouseholdColorOptionInUse(ctx, req.HouseholdID, req.UserID, req.ColorOption)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("%w: color already taken", domain.ErrConflict)
	}

	return nil
}

func validateEditHouseholdUserRequest(req HouseholdUserRequest, existing, user domain.HouseholdUser) error {
	if !auth.CanEditHouseholdUser(user, existing.User.ID) {
		return fmt.Errorf("%w: may not edit this household user", domain.ErrForbidden)
	}

	if domain.Role(req.Role) != existing.Role && !auth.CanEditRole(user) {
		return fmt.Errorf("%w: may not edit role", domain.ErrForbidden)
	}

	if req.IsActive != existing.IsActive && !auth.CanEditIsActive(user) {
		return fmt.Errorf("%w: may not edit active state", domain.ErrForbidden)
	}

	return nil
}

func (a *App) AddUserToHousehold(ctx context.Context, req HouseholdUserRequest) (domain.HouseholdUser, error) {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	err = a.validateHouseholdUserFields(ctx, req)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	hhCount, err := a.Store.HouseholdUsersCount(ctx)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	if hhCount >= domain.MaxHouseholdMembers {
		return domain.HouseholdUser{}, domain.ErrHouseholdFull
	}

	exists, err := a.HouseholdUserExists(ctx, req.UserID, req.HouseholdID)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	if exists {
		return domain.HouseholdUser{}, fmt.Errorf("%w: user already member of household", domain.ErrInvalidRequest)
	}

	now := time.Now()

	hhUser, err := a.Store.AddUserToHousehold(ctx, store.AddUserToHouseholdParams{
		HouseholdId: req.HouseholdID,
		UserId:      req.UserID,
		Role:        req.Role,
		DisplayName: req.DisplayName,
		ColorOption: req.ColorOption,
		JoinedAt:    now,
		IsActive:    req.IsActive,
	})
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	event := newHouseholdUserEvent(
		domain.HouseholdUserAdded,
		now,
		domain.ActorTypeUser,
		hhUser,
		hhUser.HouseholdID,
		hhUser.User.ID,
	)

	a.Bus.Publish(event)

	return hhUser, nil
}

func (a *App) GetHouseholdUsers(ctx context.Context) ([]domain.HouseholdUser, error) {
	hhUser, err := CheckAdmin(ctx)
	if err != nil {
		return []domain.HouseholdUser{}, err
	}

	users, err := a.Store.GetHouseholdUsers(ctx, hhUser.HouseholdID)
	if err != nil {
		return []domain.HouseholdUser{}, err
	}

	return users, nil
}

func (a *App) GetActiveHouseholdUsers(ctx context.Context) ([]domain.HouseholdUser, error) {
	hhUser, err := CheckAdmin(ctx)
	if err != nil {
		return []domain.HouseholdUser{}, err
	}

	users, err := a.Store.GetActiveHouseholdUsers(ctx, hhUser.HouseholdID)
	if err != nil {
		return []domain.HouseholdUser{}, err
	}

	return users, nil
}

func (a *App) EditHouseholdUser(ctx context.Context, req HouseholdUserRequest) (domain.HouseholdUser, error) {
	user, err := auth.AuthenticatedUser(ctx)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	existing, err := a.Store.GetHouseholdUserByID(ctx, req.UserID, req.HouseholdID)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	err = a.validateHouseholdUserFields(ctx, req)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	err = validateEditHouseholdUserRequest(req, existing, user)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	role := domain.Role(req.Role)
	if !role.IsValid() {
		return domain.HouseholdUser{}, domain.ErrInvalidRole
	}

	now := time.Now()

	updatedHouseholdUser, err := a.Store.EditHouseholdUser(ctx, store.EditHouseholdUserParams{
		Role:        role,
		DisplayName: req.DisplayName,
		ColorOption: req.ColorOption,
		IsActive:    req.IsActive,
		UserId:      req.UserID,
		HouseholdId: req.HouseholdID,
	})
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	event := newHouseholdUserEvent(domain.HouseholdUserEdited,
		now,
		domain.ActorTypeUser,
		updatedHouseholdUser,
		user.HouseholdID,
		user.User.ID,
	)

	a.Bus.Publish(event)

	return updatedHouseholdUser, nil
}

func (a *App) ActivateHouseholdUser(ctx context.Context, householdId, userId string) (domain.HouseholdUser, error) {
	reqUser, err := CheckAdmin(ctx)
	if err != nil {
		return domain.HouseholdUser{}, nil
	}

	now := time.Now()

	err = a.Store.SetHouseholdUserActive(ctx, householdId, userId)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	updatedUser, err := a.Store.GetHouseholdUserByID(ctx, userId, householdId)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	event := newHouseholdUserEvent(domain.HouseholdUserActivated,
		now,
		domain.ActorTypeUser,
		updatedUser,
		householdId,
		reqUser.User.ID,
	)

	a.Bus.Publish(event)

	return updatedUser, nil
}

func (a *App) DeactivateHouseholdUser(ctx context.Context, householdId, userId string) (domain.HouseholdUser, error) {
	reqUser, err := CheckAdmin(ctx)
	if err != nil {
		return domain.HouseholdUser{}, nil
	}

	now := time.Now()
	var updatedUser domain.HouseholdUser
	var updatedTemplates []domain.ChoreTemplate
	var updatedAssignments []domain.Assignment

	err = a.Store.WithTx(ctx, func(txStore *store.Store) error {
		templates, err := txStore.GetChoreTemplatesForUser(ctx, userId)
		if err != nil {
			return nil
		}

		for _, tmp := range templates {
			updatedTmp := tmp
			updatedTmp.Assignee = ""
			updatedTmp.UpdatedAt = now
			err := txStore.EditChoreTemplate(ctx, updatedTmp)
			if err != nil {
				return err
			}
			updatedTemplates = append(updatedTemplates, updatedTmp)
		}

		assignments, err := txStore.GetCurrentUserAssignments(ctx, userId)
		if err != nil {
			return err
		}

		for _, assignment := range assignments {
			updatedAsmt, err := txStore.EditAssignment(ctx, store.EditAssignmentParams{
				ID:             assignment.ID,
				AssignedUserID: "",
				ScheduledFor:   assignment.ScheduledFor,
				Notes:          assignment.Notes,
				Completed:      assignment.Completed,
				Canceled:       assignment.Canceled,
				UpdatedAt:      now,
				CompletedAt:    assignment.CompletedAt,
				CanceledAt:     assignment.CanceledAt,
			})
			if err != nil {
				return err
			}
			updatedAssignments = append(updatedAssignments, updatedAsmt)
		}

		err = txStore.SetHouseholdUserInactive(ctx, householdId, userId)
		if err != nil {
			return err
		}

		updatedUser, err = txStore.GetHouseholdUserByID(ctx, userId, householdId)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return domain.HouseholdUser{}, err
	}

	for _, template := range updatedTemplates {
		event := newChoreTemplateEvent(
			domain.ChoreTemplateEdited,
			now,
			domain.ActorTypeUser,
			template,
			householdId,
			reqUser.User.ID,
		)

		a.Bus.Publish(event)
	}

	for _, assignment := range updatedAssignments {
		event := newAssignmentEvent(
			domain.AssignmentEdited,
			now,
			domain.ActorTypeUser,
			assignment,
			householdId,
			reqUser.User.ID,
		)

		a.Bus.Publish(event)
	}

	event := newHouseholdUserEvent(
		domain.HouseholdUserDeactivated,
		now,
		domain.ActorTypeUser,
		updatedUser,
		householdId,
		reqUser.User.ID,
	)

	a.Bus.Publish(event)

	return updatedUser, nil
}

func (a *App) RemoveUserFromHousehold(ctx context.Context, householdId, userId string) error {
	reqUser, err := CheckAdmin(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	var userToDelete domain.HouseholdUser
	var updatedTemplates []domain.ChoreTemplate
	var updatedAssignments []domain.Assignment

	err = a.Store.WithTx(ctx, func(txStore *store.Store) error {
		templates, err := txStore.GetChoreTemplatesForUser(ctx, userId)
		if err != nil {
			return err
		}

		for _, tmp := range templates {
			updatedTmp := tmp
			updatedTmp.Assignee = ""
			updatedTmp.UpdatedAt = now
			err := txStore.EditChoreTemplate(ctx, updatedTmp)
			if err != nil {
				return err
			}

			updatedTemplates = append(updatedTemplates, updatedTmp)
		}

		assignments, err := txStore.GetCurrentUserAssignments(ctx, userId)
		if err != nil {
			return err
		}

		for _, assignment := range assignments {
			updatedAsmt, err := txStore.EditAssignment(ctx, store.EditAssignmentParams{
				ID:             assignment.ID,
				AssignedUserID: "",
				ScheduledFor:   assignment.ScheduledFor,
				Notes:          assignment.Notes,
				Completed:      assignment.Completed,
				Canceled:       assignment.Canceled,
				UpdatedAt:      now,
				CompletedAt:    assignment.CompletedAt,
				CanceledAt:     assignment.CanceledAt,
			})
			if err != nil {
				return err
			}
			updatedAssignments = append(updatedAssignments, updatedAsmt)
		}

		userToDelete, err = txStore.GetHouseholdUserByID(ctx, userId, householdId)
		if err != nil {
			return err
		}

		err = txStore.DeleteHouseholdUser(ctx, householdId, userId)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	for _, template := range updatedTemplates {
		event := newChoreTemplateEvent(
			domain.ChoreTemplateEdited,
			now,
			domain.ActorTypeUser,
			template,
			householdId,
			reqUser.User.ID,
		)

		a.Bus.Publish(event)
	}

	for _, assignment := range updatedAssignments {
		event := newAssignmentEvent(
			domain.AssignmentEdited,
			now,
			domain.ActorTypeUser,
			assignment,
			householdId,
			reqUser.User.ID,
		)

		a.Bus.Publish(event)
	}

	event := newHouseholdUserEvent(
		domain.HouseholdUserRemoved,
		now,
		domain.ActorTypeUser,
		userToDelete,
		householdId,
		reqUser.User.ID,
	)

	a.Bus.Publish(event)

	return nil
}
