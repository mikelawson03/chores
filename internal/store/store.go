package store

import (
	"database/sql"

	db "github.com/mikelawson03/chores/internal/store/db"
)

type Store struct {
	Db      *sql.DB
	Queries *db.Queries
}

func NewStore(dbConn *sql.DB) *Store {
	return &Store{
		Db:      dbConn,
		Queries: db.New(dbConn),
	}
}
