package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store/db"
)

type CreateUserParams struct {
	ID        string
	Username  string
	HashedPW  string
	FirstName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EditUserParams struct {
	ID        string
	Username  string
	FirstName string
	UpdatedAt time.Time
}

type LoginUser struct {
	User         domain.User
	PasswordHash string
}

func dbUserToDomainUser(user db.User) domain.User {
	return domain.User{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (s *Store) CreateUser(ctx context.Context, req CreateUserParams) (domain.User, error) {
	res, err := s.Queries.CreateUser(ctx, db.CreateUserParams{
		ID:           req.ID,
		Username:     req.Username,
		PasswordHash: req.HashedPW,
		FirstName:    req.FirstName,
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
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

	return domain.User{
		ID:        res.ID,
		Username:  res.Username,
		FirstName: res.FirstName,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (s *Store) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	res, err := s.Queries.GetUserByID(ctx, id)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, domain.ErrNotFound
	}

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:        res.ID,
		Username:  res.Username,
		FirstName: res.FirstName,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (s *Store) EditUser(ctx context.Context, req EditUserParams) (domain.User, error) {
	res, err := s.Queries.EditUser(ctx, db.EditUserParams{
		Username:  req.Username,
		FirstName: req.FirstName,
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
	_, err := s.Queries.DeleteUser(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.ErrNotFound
	}

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserCount(ctx context.Context) (int64, error) {
	count, err := s.Queries.GetUserCount(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *Store) GetUserWithHashedPW(ctx context.Context, username string) (LoginUser, error) {
	res, err := s.Queries.GetHashForUsername(ctx, username)
	if errors.Is(err, sql.ErrNoRows) {
		return LoginUser{}, domain.ErrNotFound
	}
	if err != nil {
		return LoginUser{}, err
	}

	user := LoginUser{
		PasswordHash: res.PasswordHash,
		User: domain.User{
			ID:        res.ID,
			Username:  res.Username,
			FirstName: res.FirstName,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		},
	}

	return user, nil
}

func (s *Store) ChangePassword(ctx context.Context, id, newPWHash string) error {
	err := s.Queries.UpdatePassword(ctx, db.UpdatePasswordParams{
		PasswordHash: newPWHash,
		ID:           id,
	})

	if err != nil {
		return err
	}

	return nil
}
