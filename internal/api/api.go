package api

import (
	"database/sql"

	"github.com/mikelawson03/chores/internal/app"
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
	"github.com/mikelawson03/chores/internal/store/db"
)

type apiCfg struct {
	App *app.App
}

func NewApiConfig(dbConn *sql.DB) apiCfg {
	return apiCfg{
		App: &app.App{
			Store: &store.Store{
				Db:      dbConn,
				Queries: db.New(dbConn),
			},
			Assignments: make([]domain.Assignment, 0),
		},
	}
}
