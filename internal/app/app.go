package app

import (
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
)

type App struct {
	Store       *store.Store
	Templates   map[string]domain.ChoreTemplate
	Assignments []domain.Assignment
	Users       []domain.User
}
