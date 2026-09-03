package app

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mikelawson03/chores/internal/auth"
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
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

func CheckAdmin(ctx context.Context) (domain.HouseholdUser, error) {
	user, err := auth.AuthenticatedUser(ctx)
	if err != nil {
		return domain.HouseholdUser{}, err
	}

	if user.Role != domain.RoleAdmin {
		return domain.HouseholdUser{}, domain.ErrForbidden
	}

	return user, nil
}

func (a *App) Bootstrap(ctx context.Context, username, firstName, password string) (domain.User, error) {
	userCount, err := a.Store.GetUserCount(ctx)
	if err != nil {
		return domain.User{}, err
	}

	if userCount != 0 {
		return domain.User{}, errors.New("may only use bootstrap if no users exist")
	}

	if err = validatePassword(password); err != nil {
		return domain.User{}, err
	}

	hashedPW, err := hashPassword(password)
	if err != nil {
		return domain.User{}, err
	}

	user, err := a.Store.CreateUser(ctx, store.CreateUserParams{
		ID:        uuid.NewString(),
		Username:  username,
		HashedPW:  hashedPW,
		FirstName: firstName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
