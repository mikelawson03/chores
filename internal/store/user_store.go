package store

import (
	"context"
	"fmt"
	"time"

	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store/db"
)

type CreateUserParams struct {
	ID        string
	Username  string
	Role      domain.Role
	FirstName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EditUserParams struct {
	ID        string
	Username  string
	Role      domain.Role
	FirstName string
	UpdatedAt time.Time
}

func dbUserToDomainUser(user db.User) domain.User {
	return domain.User{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		Role:      domain.Role(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (s *Store) CreateUser(ctx context.Context, req CreateUserParams) (domain.User, error) {
	res, err := s.Queries.CreateUser(ctx, db.CreateUserParams{
		ID:        req.ID,
		Username:  req.Username,
		Role:      string(req.Role),
		FirstName: req.FirstName,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	})

	if err != nil {
		return domain.User{}, err
	}

	user := dbUserToDomainUser(res)

	return user, nil
}

func (s *Store) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	res, err := s.Queries.GetUserByUsername(ctx, username)
	if err != nil {
		return domain.User{}, err
	}

	return dbUserToDomainUser(res), nil
}

func (s *Store) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	res, err := s.Queries.GetUserByID(ctx, id)

	if err != nil {
		return domain.User{}, err
	}

	return dbUserToDomainUser(res), nil
}

func (s *Store) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	res, err := s.Queries.GetAllUsers(ctx)

	if err != nil {
		return []domain.User{}, err
	}

	for _, user := range res {
		users = append(users, dbUserToDomainUser(user))
	}

	return users, nil
}

func (s *Store) EditUser(ctx context.Context, req EditUserParams) (domain.User, error) {
	res, err := s.Queries.EditUser(ctx, db.EditUserParams{
		Username:  req.Username,
		FirstName: req.FirstName,
		Role:      string(req.Role),
		UpdatedAt: req.UpdatedAt,
		ID:        req.ID,
	})

	if err != nil {
		return domain.User{}, err
	}

	user := dbUserToDomainUser(res)

	return user, nil
}

func (s *Store) DeleteUser(ctx context.Context, id string) error {
	res, err := s.Queries.DeleteUser(ctx, id)
	fmt.Println("ID: ", res)
	if err != nil {
		return err
	}

	return nil
}
