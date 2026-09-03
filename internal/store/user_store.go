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

type EditHouseholdUserParams struct {
	Role        domain.Role
	DisplayName string
	ColorOption int
	IsActive    bool
	UserId      string
	HouseholdId string
}

type AddUserToHouseholdParams struct {
	HouseholdId string
	UserId      string
	Role        string
	DisplayName string
	ColorOption int
	JoinedAt    time.Time
	IsActive    bool
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

func mapGetHouseholdUsersRow(user db.GetHouseholdUsersRow) (domain.HouseholdUser, error) {

	role := domain.Role(user.Role)
	if !role.IsValid() {
		return domain.HouseholdUser{}, domain.ErrInvalidRole
	}

	displayName := user.FirstName
	if user.DisplayName.Valid {
		displayName = user.DisplayName.String
	}

	return domain.HouseholdUser{
		HouseholdID: user.HouseholdID,
		Role:        role,
		DisplayName: displayName,
		ColorOption: int(user.ColorOption),
		JoinedAt:    user.JoinedAt,
		IsActive:    user.IsActive,
		User: domain.User{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}

func mapGetHouseholdUserByIdRow(user db.GetHouseholdUserByIDRow) (domain.HouseholdUser, error) {
	role := domain.Role(user.Role)
	if !role.IsValid() {
		return domain.HouseholdUser{}, domain.ErrInvalidRole
	}

	displayName := user.FirstName
	if user.DisplayName.Valid {
		displayName = user.DisplayName.String
	}

	return domain.HouseholdUser{
		HouseholdID: user.HouseholdID,
		Role:        role,
		DisplayName: displayName,
		ColorOption: int(user.ColorOption),
		JoinedAt:    user.JoinedAt,
		IsActive:    user.IsActive,
		User: domain.User{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
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

func (s *Store) GetHouseholdUsers(ctx context.Context, householdId string) ([]domain.HouseholdUser, error) {
	var householdUsers []domain.HouseholdUser

	res, err := s.Queries.GetHouseholdUsers(ctx, householdId)

	if err != nil {
		return []domain.HouseholdUser{}, err
	}

	for _, user := range res {

		hhUser, err := mapGetHouseholdUsersRow(user)
		if err != nil {
			return []domain.HouseholdUser{}, err
		}

		householdUsers = append(householdUsers, hhUser)
	}

	return householdUsers, nil
}

func (s *Store) GetHouseholdUserByID(ctx context.Context, householdId, userId string) (domain.HouseholdUser, error) {
	res, err := s.Queries.GetHouseholdUserByID(ctx, db.GetHouseholdUserByIDParams{
		HouseholdID: householdId,
		UserID:      userId,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return domain.HouseholdUser{}, domain.ErrNotFound
	}

	if err != nil {
		return domain.HouseholdUser{}, err
	}

	hhUser, err := mapGetHouseholdUserByIdRow(res)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	return hhUser, nil
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

func (s *Store) AddUserToHousehold(ctx context.Context, req AddUserToHouseholdParams) (domain.HouseholdUser, error) {
	err := s.Queries.AddUserToHousehold(ctx, db.AddUserToHouseholdParams{
		HouseholdID: req.HouseholdId,
		UserID:      req.UserId,
		Role:        string(req.Role),
		DisplayName: stringToNullString(req.DisplayName),
		ColorOption: int64(req.ColorOption),
		JoinedAt:    req.JoinedAt,
		IsActive:    req.IsActive,
	})

	if err != nil {
		return domain.HouseholdUser{}, err
	}

	hhUser, err := s.GetHouseholdUserByID(ctx, req.HouseholdId, req.UserId)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	return hhUser, nil
}

func (s *Store) EditHouseholdUser(ctx context.Context, req EditHouseholdUserParams) (domain.HouseholdUser, error) {
	_, err := s.Queries.EditHouseholdUser(ctx, db.EditHouseholdUserParams{
		Role:        string(req.Role),
		DisplayName: stringToNullString(req.DisplayName),
		ColorOption: int64(req.ColorOption),
		IsActive:    req.IsActive,
		UserID:      req.UserId,
		HouseholdID: req.HouseholdId,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return domain.HouseholdUser{}, domain.ErrNotFound
	}

	if err != nil {
		return domain.HouseholdUser{}, err
	}

	hhUser, err := s.GetHouseholdUserByID(ctx, req.HouseholdId, req.UserId)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	return hhUser, nil
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

func (s *Store) HouseholdColorOptionInUse(ctx context.Context, householdId, userId string, colorOption int) (bool, error) {
	return s.Queries.HouseholdColorOptionInUse(ctx, db.HouseholdColorOptionInUseParams{
		HouseholdID: householdId,
		ColorOption: int64(colorOption),
		UserID:      userId,
	})
}

func (s *Store) HouseholdUsersCount(ctx context.Context) (int64, error) {
	return s.Queries.HouseholdUsersCount(ctx)
}
