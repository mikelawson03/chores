package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
	"github.com/mikelawson03/chores/internal/auth"
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
)

type LoginResult struct {
	Token string
}

func newUserEvent(eventType domain.UserEventType,
	occurredAt time.Time,
	actorType domain.ActorType,
	user domain.User,
	householdID,
	userID string) domain.Event[domain.User] {
	return domain.Event[domain.User]{
		Type:        string(eventType),
		OccurredAt:  occurredAt,
		HouseholdID: householdID,
		Actor: domain.Actor{
			Type: actorType,
			ID:   userID,
		},
		Payload: user,
	}
}

func hashPassword(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

func validatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return domain.ErrPasswordRequired
	}

	if len(password) < 8 {
		return domain.ErrPasswordTooShort
	}

	return nil
}

func (a *App) UsernameExists(ctx context.Context, username string) (bool, error) {
	_, err := a.Store.GetUserByUsername(ctx, username)

	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil

}

func (a *App) validateNewUserRequest(ctx context.Context, username, firstName, password string) error {
	if username == "" {
		return fmt.Errorf("%w: username required", domain.ErrInvalidRequest)
	}

	exists, err := a.UsernameExists(ctx, username)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("%w: username already exists", domain.ErrInvalidRequest)
	}

	if firstName == "" {
		return fmt.Errorf("%w: must provide first name", domain.ErrInvalidRequest)
	}

	if err = validatePassword(password); err != nil {
		return err
	}

	return nil
}

func (a *App) validateEditUserRequest(ctx context.Context, newUsername, firstName, reqUserID string, existing domain.User) error {
	if !auth.CanEditUser(reqUserID, existing.ID) {
		return fmt.Errorf("%w: may not edit user", domain.ErrForbidden)
	}

	userWithName, err := a.Store.GetUserByUsername(ctx, newUsername)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if err == nil && userWithName.ID != existing.ID {
		return fmt.Errorf("%w: username already in use", domain.ErrInvalidRequest)
	}

	if firstName == "" {
		return fmt.Errorf("%w: must provide first name", domain.ErrInvalidRequest)
	}

	return nil
}

func (a *App) validateDeleteUserRequest(ctx context.Context, reqUserID, userID string) error {
	if !auth.CanDeleteUser(reqUserID, userID) {
		return fmt.Errorf("%w: may not delete user", domain.ErrForbidden)
	}

	hasMemberships, err := a.Store.UserHasHouseholdMemberships(ctx, userID)
	if err != nil {
		return err
	}

	if hasMemberships {
		return domain.ErrHasHouseholdMemberships
	}

	hasAssignments, err := a.Store.UserHasAssignments(ctx, userID)
	if err != nil {
		return err
	}

	if hasAssignments {
		return domain.ErrHasAssignments
	}

	return nil
}

func (a *App) CreateNewUser(ctx context.Context, username, firstName, password string) (domain.User, error) {
	err := a.validateNewUserRequest(ctx, username, firstName, password)
	if err != nil {
		return domain.User{}, err
	}

	hashedPW, err := hashPassword(password)
	if err != nil {
		return domain.User{}, err
	}

	now := time.Now()

	user, err := a.Store.CreateUser(ctx, store.CreateUserParams{
		ID:        uuid.NewString(),
		Username:  username,
		HashedPW:  hashedPW,
		FirstName: firstName,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (a *App) GetUserByID(ctx context.Context, userID string) (domain.User, error) {
	reqUserID, ok := auth.UserIDFromContext(ctx)
	if !ok {
		return domain.User{}, fmt.Errorf("%w: authenticated user ID missing from context", domain.ErrUnauthorized)
	}

	if !auth.CanGetUser(reqUserID, userID) {
		return domain.User{}, domain.ErrForbidden
	}

	user, err := a.Store.GetUserByID(ctx, userID)
	if errors.Is(err, domain.ErrNotFound) {
		return domain.User{}, fmt.Errorf("%w: user", err)
	}
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (a *App) EditUser(ctx context.Context, id, newUsername, firstName string) (domain.User, error) {
	reqUserID, ok := auth.UserIDFromContext(ctx)
	if !ok {
		return domain.User{}, fmt.Errorf("%w: authenticated user ID missing from context", domain.ErrUnauthorized)
	}

	existing, err := a.Store.GetUserByID(ctx, id)
	if errors.Is(err, domain.ErrNotFound) {
		return domain.User{}, fmt.Errorf("%w: user", err)
	}
	if err != nil {
		return domain.User{}, err
	}

	err = a.validateEditUserRequest(ctx, newUsername, firstName, reqUserID, existing)
	if err != nil {
		return domain.User{}, err
	}

	now := time.Now()

	updatedUser, err := a.Store.EditUser(ctx, store.EditUserParams{
		ID:        id,
		Username:  newUsername,
		FirstName: firstName,
		UpdatedAt: now,
	})
	if err != nil {
		return domain.User{}, err
	}

	// TODO: When multiple households are supported, publish this event
	// once for each household the user belongs to.
	event := newUserEvent(domain.UserEdited,
		now,
		domain.ActorTypeUser,
		updatedUser,
		domain.DefaultHouseholdID,
		reqUserID,
	)

	a.Bus.Publish(event)

	return updatedUser, nil
}

func (a *App) DeleteUser(ctx context.Context, id string) error {
	reqUserID, ok := auth.UserIDFromContext(ctx)
	if !ok {
		return fmt.Errorf("%w: authenticated user ID missing from context", domain.ErrUnauthorized)
	}

	err := a.validateDeleteUserRequest(ctx, reqUserID, id)
	if err != nil {
		return err
	}

	err = a.Store.DeleteUser(ctx, id)
	if errors.Is(err, domain.ErrNotFound) {
		return fmt.Errorf("%w: user", err)
	}
	if err != nil {
		return err
	}

	return nil
}

func (a *App) LoginUser(ctx context.Context, username, password string) (LoginResult, error) {

	user, err := a.Store.GetUserWithHashedPW(ctx, username)
	if err != nil {
		return LoginResult{}, err
	}

	match, err := argon2id.ComparePasswordAndHash(password, user.PasswordHash)
	if err != nil {
		return LoginResult{}, err
	}

	if !match {
		return LoginResult{}, domain.ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(user.User.ID, a.Config.JWTSigninSecret, time.Hour*24*30)
	if err != nil {
		return LoginResult{}, err
	}

	return LoginResult{
		Token: token,
	}, err

}

func (a *App) ChangePassword(ctx context.Context, oldPassword, newPassword string) error {
	reqUserID, ok := auth.UserIDFromContext(ctx)
	if !ok {
		return fmt.Errorf("%w: authenticated user ID missing from context", domain.ErrUnauthorized)
	}

	userWithHash, err := a.Store.GetUserWithHashedPWByID(ctx, reqUserID)
	if err != nil {
		return err
	}

	match, err := argon2id.ComparePasswordAndHash(oldPassword, userWithHash.PasswordHash)
	if err != nil {
		return err
	}

	if !match {
		return domain.ErrInvalidCredentials
	}

	err = validatePassword(newPassword)
	if err != nil {
		return err
	}

	newPWHash, err := hashPassword(newPassword)
	if err != nil {
		return err
	}

	now := time.Now()

	err = a.Store.ChangePassword(ctx, reqUserID, newPWHash, now)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) ResetPassword(ctx context.Context, userID, tempPassword string) error {
	reqUser, err := CheckAdmin(ctx)
	if err != nil {
		return err
	}

	_, err = a.Store.GetHouseholdUserByID(ctx, userID, reqUser.HouseholdID)
	if errors.Is(err, domain.ErrNotFound) {
		return domain.ErrForbidden
	}
	if err != nil {
		return err
	}

	tempPWHash, err := hashPassword(tempPassword)
	if err != nil {
		return err
	}

	now := time.Now()

	err = a.Store.ChangePassword(ctx, userID, tempPWHash, now)
	if err != nil {
		return err
	}

	return nil
}
