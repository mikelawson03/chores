package app

import (
	"context"
	"errors"
	"time"

	"github.com/mikelawson03/chores/internal/auth"
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
)

var (
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrNotFound           = errors.New("not found")
	ErrValidation         = errors.New("invalid request")
	ErrInvalidRole        = errors.New("invalid role")
)

type Config struct {
	JWTSigninSecret  string
	DevLoginPassword string
	DevUsername      string
}

type App struct {
	Store  *store.Store
	Config Config
}

func NewConfig(secret, password, user string) Config {
	return Config{
		JWTSigninSecret:  secret,
		DevLoginPassword: password,
		DevUsername:      user,
	}
}

func NewApp(store *store.Store, config Config) *App {
	return &App{
		Store:  store,
		Config: config,
	}
}

func (a *App) getHorizonWindow() (start time.Time, end time.Time) {
	now := time.Now()
	current_day := now.Weekday()
	daysSinceWeekStart := (int(current_day) + 6) % 7
	startOfWeek := now.AddDate(0, 0, -daysSinceWeekStart)
	horizonStart := time.Date(
		startOfWeek.Year(),
		startOfWeek.Month(),
		startOfWeek.Day(),
		0, 0, 0, 0,
		startOfWeek.Location(),
	)
	horizonEnd := horizonStart.AddDate(0, 0, 28)

	return horizonStart, horizonEnd
}

func (a *App) getMonthlyPlanningEnd(horizonStart, horizonEnd time.Time) time.Time {
	thisMonthStart := time.Date(horizonStart.Year(), horizonStart.Month(), 1, 0, 0, 0, 0, horizonStart.Location())
	nextMonthStart := thisMonthStart.AddDate(0, 1, 0)
	if nextMonthStart.After(horizonEnd) {
		return nextMonthStart
	}
	return nextMonthStart.AddDate(0, 1, 0)
}

func AuthenticatedUser(ctx context.Context) (domain.User, error) {
	user, ok := auth.UserFromContext(ctx)
	if !ok {
		return domain.User{}, ErrUnauthorized
	}

	return user, nil
}

func CheckAdmin(ctx context.Context) (domain.User, error) {
	user, err := AuthenticatedUser(ctx)
	if err != nil {
		return domain.User{}, err
	}

	if user.Role != domain.RoleAdmin {
		return domain.User{}, ErrForbidden
	}

	return user, nil
}
