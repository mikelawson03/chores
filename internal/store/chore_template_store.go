package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store/db"
)

func dbTemplateToDomainTemplate(dbTmp db.ChoreTemplate) domain.ChoreTemplate {
	return domain.ChoreTemplate{
		ID:           dbTmp.ID,
		Name:         dbTmp.Name,
		Cadence:      domain.Cadence(dbTmp.Cadence),
		Assignee:     dbTmp.Assignee.String,
		Instructions: dbTmp.Instructions.String,
		Duration:     int(dbTmp.Duration),
		CreatedAt:    dbTmp.CreatedAt,
		UpdatedAt:    dbTmp.UpdatedAt,
	}
}

func (s *Store) AddChoreTemplate(ctx context.Context, tmp domain.ChoreTemplate) error {
	var assignee sql.NullString
	var instructions sql.NullString

	if tmp.Assignee != "" {
		assignee = sql.NullString{
			String: tmp.Assignee,
			Valid:  true,
		}
	} else {
		assignee.Valid = false
	}

	if tmp.Instructions != "" {
		instructions = sql.NullString{
			String: tmp.Instructions,
			Valid:  true,
		}
	} else {
		instructions.Valid = false
	}

	err := s.Queries.CreateChoreTemplate(ctx, db.CreateChoreTemplateParams{
		ID:           tmp.ID,
		Name:         tmp.Name,
		Cadence:      string(tmp.Cadence),
		Assignee:     assignee,
		Instructions: instructions,
		Duration:     int64(tmp.Duration),
		CreatedAt:    tmp.CreatedAt,
		UpdatedAt:    tmp.UpdatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetTemplateByName(ctx context.Context, name string) (domain.ChoreTemplate, error) {
	dbTmp, err := s.Queries.GetChoreTemplateByName(ctx, name)
	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	tmp := dbTemplateToDomainTemplate(dbTmp)

	return tmp, nil
}

func (s *Store) GetTemplateByID(ctx context.Context, id string) (domain.ChoreTemplate, error) {
	dbTmp, err := s.Queries.GetChoreTemplateByID(ctx, id)

	if err != nil {
		return domain.ChoreTemplate{}, err
	}

	tmp := dbTemplateToDomainTemplate(dbTmp)

	return tmp, nil
}

func (s *Store) GetChoreTemplates(ctx context.Context) ([]domain.ChoreTemplate, error) {
	domainTmps := make([]domain.ChoreTemplate, 0)

	dbTmps, err := s.Queries.GetAllChoreTemplates(ctx)
	if err != nil {
		return []domain.ChoreTemplate{}, nil
	}

	for _, tmp := range dbTmps {
		domainTmps = append(domainTmps, dbTemplateToDomainTemplate(tmp))
	}

	return domainTmps, err
}

func (s *Store) EditChoreTemplate(ctx context.Context, tmp domain.ChoreTemplate) error {
	var assignee sql.NullString
	var instructions sql.NullString

	if tmp.Assignee != "" {
		assignee = sql.NullString{
			String: tmp.Assignee,
			Valid:  true,
		}
	} else {
		assignee.Valid = false
	}

	if tmp.Instructions != "" {
		instructions = sql.NullString{
			String: tmp.Instructions,
			Valid:  true,
		}
	} else {
		instructions.Valid = false
	}

	err := s.Queries.EditChoreTemplate(ctx, db.EditChoreTemplateParams{
		Name:         tmp.Name,
		Cadence:      string(tmp.Cadence),
		Assignee:     assignee,
		Instructions: instructions,
		Duration:     int64(tmp.Duration),
		UpdatedAt:    time.Now(),
		ID:           tmp.ID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteChoreTemplate(ctx context.Context, id string) error {
	err := s.Queries.DeleteChoreTemplate(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
