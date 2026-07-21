package auth

import (
	"context"

	"github.com/mikelawson03/chores/internal/domain"
)

type contextKey string

const userContextKey contextKey = "user"

func WithUser(ctx context.Context, user domain.User) context.Context {

	return context.WithValue(ctx, userContextKey, user)
}

func UserFromContext(ctx context.Context) (domain.User, bool) {
	user, ok := ctx.Value(userContextKey).(domain.User)
	return user, ok
}
