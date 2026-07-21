package store

import (
	"context"

	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store/db"
)

func dbUserToDomainUser(user db.User) domain.User {
	return domain.User{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (s *Store) CreateUser(ctx context.Context, u domain.User) error {
	err := s.Queries.CreateUser(ctx, db.CreateUserParams{
		ID:        u.ID,
		Username:  u.Username,
		Role:      u.Role,
		FirstName: u.FirstName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	})

	if err != nil {
		return err
	}

	return nil
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

func (s *Store) EditUser(ctx context.Context, u domain.User) error {
	err := s.Queries.EditUser(ctx, db.EditUserParams{
		Username:  u.Username,
		UpdatedAt: u.UpdatedAt,
		ID:        u.ID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteUser(ctx context.Context, id string) error {
	err := s.Queries.DeleteUser(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
