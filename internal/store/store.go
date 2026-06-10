package store

import (
	"database/sql"

	db "github.com/mikelawson03/chores/internal/store/db"
)

type Store struct {
	Db      *sql.DB
	Queries *db.Queries
}
