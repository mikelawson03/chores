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

func (a *App) CreateNewUser(ctx context.Context, username, role, firstName string) (domain.User, error) {
	exists, err := a.UsernameExists(ctx, username)

	if err != nil {
		return domain.User{}, err
	}

	if exists {
		return domain.User{}, fmt.Errorf("User with name %s already exists", username)
	}

	user := domain.User{
		ID:        uuid.NewString(),
		Username:  username,
		Role:      role,
		FirstName: firstName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = a.Store.CreateUser(context.Background(), user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (a *App) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	users, err := a.Store.GetAllUsers(ctx)
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (a *App) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	user, err := a.Store.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (a *App) EditUser(ctx context.Context, id, newUsername, requesterID string) (domain.User, error) {
	existing, err := a.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	if id != requesterID {
		return domain.User{}, errors.New("May not edit other users")
	}

	userWithName, err := a.Store.GetUserByUsername(ctx, newUsername)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, err
	}

	if err == nil && userWithName.ID != requesterID {
		return domain.User{}, fmt.Errorf("User with name %s already exists", newUsername)
	}

	updatedUser := domain.User{
		ID:        existing.ID,
		Username:  newUsername,
		Role:      existing.Role,
		CreatedAt: existing.CreatedAt,
		UpdatedAt: time.Now(),
	}

	err = a.Store.EditUser(ctx, updatedUser)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (a *App) DeleteUser(ctx context.Context, id, requesterID string) error {
	_, err := a.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	if id != requesterID {
		return fmt.Errorf("May not delete other users")
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
