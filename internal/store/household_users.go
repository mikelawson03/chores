package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store/db"
)

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

func mapGetActiveHouseholdUsersRow(user db.GetActiveHouseholdUsersRow) (domain.HouseholdUser, error) {
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

func (s *Store) GetActiveHouseholdUsers(ctx context.Context, householdId string) ([]domain.HouseholdUser, error) {
	var activeHouseholdUsers []domain.HouseholdUser

	res, err := s.Queries.GetActiveHouseholdUsers(ctx, householdId)
	if err != nil {
		return []domain.HouseholdUser{}, err
	}

	for _, user := range res {
		hhUser, err := mapGetActiveHouseholdUsersRow(user)
		if err != nil {
			return []domain.HouseholdUser{}, err
		}

		activeHouseholdUsers = append(activeHouseholdUsers, hhUser)
	}

	return activeHouseholdUsers, nil
}

func (s *Store) GetHouseholdUserByID(ctx context.Context, userId, householdId string) (domain.HouseholdUser, error) {
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

	hhUser, err := s.GetHouseholdUserByID(ctx, req.UserId, req.HouseholdId)
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

	hhUser, err := s.GetHouseholdUserByID(ctx, req.UserId, req.HouseholdId)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	return hhUser, nil
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

func (s *Store) SetHouseholdUserActive(ctx context.Context, householdId, userId string) error {
	return s.Queries.SetHouseholdUserActive(ctx, db.SetHouseholdUserActiveParams{
		HouseholdID: householdId,
		UserID:      userId,
	})
}

func (s *Store) SetHouseholdUserInactive(ctx context.Context, householdId, userId string) error {
	return s.Queries.SetHouseholdUserInactive(ctx, db.SetHouseholdUserInactiveParams{
		HouseholdID: householdId,
		UserID:      userId,
	})
}

func (s *Store) DeleteHouseholdUser(ctx context.Context, householdId, userId string) error {
	result, err := s.Queries.DeleteHouseholdUser(ctx, db.DeleteHouseholdUserParams{
		HouseholdID: householdId,
		UserID:      userId,
	})

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return fmt.Errorf(
			"%w: expected to remove 1 household user, removed %d",
			domain.ErrNotFound,
			rows,
		)
	}

	return nil
}
