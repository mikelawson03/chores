package store

import (
	"context"

	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store/db"
)

func (s *Store) CreateTemplate(ctx context.Context, tmp domain.ChoreTemplate) error {
	err := s.Queries.CreateTemplate(ctx, db.CreateTemplateParams{
		ID:        tmp.ID,
		Name:      tmp.Name,
		Cadence:   tmp.Cadence,
		Shared:    tmp.Shared,
		Assignee:  tmp.Assignee,
		Duration:  int64(tmp.Duration),
		CreatedAt: tmp.CreatedAt,
		UpdatedAt: tmp.UpdatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}
