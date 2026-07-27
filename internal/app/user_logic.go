package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/mikelawson03/chores/internal/auth"
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
)

type LoginResult struct {
	User  domain.User
	Token string
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

func (a *App) validateNewUserRequest(ctx context.Context, username, role, firstName string) error {
	exists, err := a.UsernameExists(ctx, username)
	if err != nil {
		return err
	}

	if exists {
		err = fmt.Errorf("%w: username already exists", ErrValidation)
	}

	userRole := domain.Role(role)
	if !userRole.IsValid() {
		return ErrInvalidRole
	}

	if firstName == "" {
		return errors.New("must provide first name")
	}

	return nil
}

func (a *App) validateEditUserRequest(ctx context.Context, newUsername, role, firstName string, user domain.User) error {
	userRole := domain.Role(role)
	if !userRole.IsValid() {
		return ErrInvalidRole
	}

	userWithName, err := a.Store.GetUserByUsername(ctx, newUsername)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if err == nil && userWithName.ID != user.ID {
		return errors.New("username already in use")
	}

	if firstName == "" {
		return errors.New("must provide first name")
	}

	return nil
}

func (a *App) CreateNewUser(ctx context.Context, username, role, firstName string) (domain.User, error) {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return domain.User{}, err
	}

	err = a.validateNewUserRequest(ctx, username, role, firstName)
	if err != nil {
		return domain.User{}, err
	}

	user, err := a.Store.CreateUser(ctx, store.CreateUserParams{
		ID:        uuid.NewString(),
		Username:  username,
		Role:      domain.Role(role),
		FirstName: firstName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (a *App) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return []domain.User{}, err
	}

	users, err := a.Store.GetAllUsers(ctx)
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (a *App) GetUserByID(ctx context.Context, id string) (domain.User, error) {

	reqUser, err := AuthenticatedUser(ctx)
	if err != nil {
		return domain.User{}, err
	}

	if !auth.CanGetUser(reqUser, id) {
		return domain.User{}, ErrForbidden
	}

	user, err := a.Store.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (a *App) EditUser(ctx context.Context, id, newUsername, role, firstName string) (domain.User, error) {
	user, err := AuthenticatedUser(ctx)
	if err != nil {
		return domain.User{}, err
	}

	existing, err := a.Store.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	err = a.validateEditUserRequest(ctx, newUsername, role, firstName, user)

	if newUsername != existing.Username && !auth.CanEditUserName(user, existing.ID) {
		return domain.User{}, ErrForbidden
	}

	userRole := domain.Role(role)

	if userRole != existing.Role && !auth.CanEditRole(user) {
		return domain.User{}, ErrForbidden
	}

	if firstName != existing.FirstName && !auth.CanEditFirstName(user, existing.ID) {
		return domain.User{}, ErrForbidden
	}

	updatedUser, err := a.Store.EditUser(ctx, store.EditUserParams{
		ID:        id,
		Username:  newUsername,
		Role:      userRole,
		FirstName: firstName,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (a *App) DeleteUser(ctx context.Context, id string) error {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return err
	}

	err = a.Store.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) LoginUser(ctx context.Context, username, password string) (LoginResult, error) {

	devUsername := os.Getenv("DEV_USERNAME")
	devPassword := os.Getenv("DEV_PASSWORD")

	if username != devUsername || password != devPassword {
		return LoginResult{}, errors.New("Invalid credentials")
	}

	user, err := a.Store.GetUserByUsername(ctx, username)
	if err != nil {
		return LoginResult{}, err
	}

	token, err := auth.GenerateToken(user.ID, a.Config.JWTSigninSecret, time.Hour*24*30)
	if err != nil {
		return LoginResult{}, err
	}

	return LoginResult{
		User:  user,
		Token: token,
	}, err

}
