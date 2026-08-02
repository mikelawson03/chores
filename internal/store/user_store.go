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
	HashedPW  string
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

type LoginUser struct {
	User         domain.User
	PasswordHash string
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
		ID:           req.ID,
		Username:     req.Username,
		PasswordHash: req.HashedPW,
		Role:         string(req.Role),
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

	role := domain.Role(res.Role)
	if !role.IsValid() {
		return domain.User{}, domain.ErrInvalidRole
	}

	return domain.User{
		ID:        res.ID,
		Username:  res.Username,
		FirstName: res.FirstName,
		Role:      domain.Role(res.Role),
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (s *Store) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	res, err := s.Queries.GetUserByID(ctx, id)

	if err != nil {
		return domain.User{}, err
	}

	role := domain.Role(res.Role)
	if !role.IsValid() {
		return domain.User{}, domain.ErrInvalidRole
	}

	return domain.User{
		ID:        res.ID,
		Username:  res.Username,
		FirstName: res.FirstName,
		Role:      domain.Role(res.Role),
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (s *Store) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	res, err := s.Queries.GetAllUsers(ctx)

	if err != nil {
		return []domain.User{}, err
	}

	for _, user := range res {
		role := domain.Role(user.Role)
		if !role.IsValid() {
			return []domain.User{}, domain.ErrInvalidRole
		}
		users = append(users, domain.User{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			Role:      role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
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

func (s *Store) GetUserCount(ctx context.Context) (int64, error) {
	count, err := s.Queries.GetUserCount(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *Store) GetUserWithHashedPW(ctx context.Context, username string) (LoginUser, error) {
	res, err := s.Queries.GetHashForUsername(ctx, username)
	if err != nil {
		return LoginUser{}, err
	}

	role := domain.Role(res.Role)
	if !role.IsValid() {
		return LoginUser{}, domain.ErrInvalidRole
	}

	user := LoginUser{
		PasswordHash: res.PasswordHash,
		User: domain.User{
			ID:        res.ID,
			Username:  res.Username,
			FirstName: res.FirstName,
			Role:      role,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		},
	}

	return user, nil
}
