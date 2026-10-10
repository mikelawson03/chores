package auth

import (
	"context"

	"github.com/mikelawson03/chores/internal/domain"
)

type contextKey string

const (
	userIdContextKey        contextKey = "userId"
	householdUserContextKey contextKey = "householdUser"
)

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIdContextKey, userID)
}

func WithHouseholdUser(ctx context.Context, householdUser domain.HouseholdUser) context.Context {

	return context.WithValue(ctx, householdUserContextKey, householdUser)
}

func UserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIdContextKey).(string)
	return userID, ok
}

func HouseholdUserFromContext(ctx context.Context) (domain.HouseholdUser, bool) {
	householdUser, ok := ctx.Value(householdUserContextKey).(domain.HouseholdUser)
	return householdUser, ok
}

func AuthenticatedUser(ctx context.Context) (domain.HouseholdUser, error) {
	user, ok := HouseholdUserFromContext(ctx)
	if !ok {
		return domain.HouseholdUser{}, domain.ErrUnauthorized
	}

	return user, nil
}
