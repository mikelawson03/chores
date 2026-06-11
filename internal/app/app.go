package app

import (
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
)

type App struct {
	Store       *store.Store
	Assignments []domain.Assignment
}
