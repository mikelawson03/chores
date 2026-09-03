package auth

import (
	"context"

	"github.com/mikelawson03/chores/internal/domain"
)

type contextKey string

const userContextKey contextKey = "user"

func WithUser(ctx context.Context, user domain.HouseholdUser) context.Context {

	return context.WithValue(ctx, userContextKey, user)
}

func UserFromContext(ctx context.Context) (domain.HouseholdUser, bool) {
	user, ok := ctx.Value(userContextKey).(domain.HouseholdUser)
	return user, ok
}

func AuthenticatedUser(ctx context.Context) (domain.HouseholdUser, error) {
	user, ok := UserFromContext(ctx)
	if !ok {
		return domain.HouseholdUser{}, domain.ErrUnauthorized
	}

	return user, nil
}
