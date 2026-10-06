package store

import (
	"context"
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

func (s *Store) WithTx(
	ctx context.Context,
	fn func(*Store) error,
) error {
	tx, err := s.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	txStore := &Store{
		Queries: s.Queries.WithTx(tx),
	}

	if err := fn(txStore); err != nil {
		return err
	}

	return tx.Commit()
}
