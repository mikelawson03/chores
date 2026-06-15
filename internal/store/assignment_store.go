package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store/db"
)

func dbAssignmentToDomainAssignment(dbAssignment db.Assignment) domain.Assignment {
	var completedAt time.Time
	var canceledAt time.Time

	if dbAssignment.CompletedAt.Valid {
		completedAt = dbAssignment.CompletedAt.Time
	}

	if dbAssignment.CanceledAt.Valid {
		canceledAt = dbAssignment.CanceledAt.Time
	}

	return domain.Assignment{
		ID:             dbAssignment.ID,
		TemplateID:     dbAssignment.TemplateID,
		AssignedUserID: dbAssignment.AssignedUserID,
		ScheduledFor:   dbAssignment.ScheduledFor,
		Completed:      dbAssignment.Completed,
		Canceled:       dbAssignment.Canceled,
		CreatedAt:      dbAssignment.CreatedAt,
		UpdatedAt:      dbAssignment.UpdatedAt,
		CompletedAt:    completedAt,
		CanceledAt:     canceledAt,
	}
}

func (s *Store) AddAssignment(ctx context.Context, assignment domain.Assignment) error {

	err := s.Queries.CreateAssignment(ctx, db.CreateAssignmentParams{
		ID:             assignment.ID,
		TemplateID:     assignment.TemplateID,
		AssignedUserID: assignment.AssignedUserID,
		ScheduledFor:   assignment.ScheduledFor,
		Completed:      false,
		Canceled:       false,
		CreatedAt:      assignment.CreatedAt,
		UpdatedAt:      assignment.UpdatedAt,
		CompletedAt:    sql.NullTime{Valid: false},
		CanceledAt:     sql.NullTime{Valid: false},
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetAssignmentByID(ctx context.Context, id string) (domain.Assignment, error) {
	dbAssignment, err := s.Queries.GetAssignmentByID(ctx, id)
	if err != nil {
		return domain.Assignment{}, err
	}

	assignment := dbAssignmentToDomainAssignment(dbAssignment)

	return assignment, nil

}

func (s *Store) GetAllAssignments(ctx context.Context) ([]domain.Assignment, error) {
	dbAssignments, err := s.Queries.GetAllAssignments(ctx)
	if err != nil {
		return []domain.Assignment{}, err
	}

	assignments := []domain.Assignment{}

	for _, assignment := range dbAssignments {
		assignments = append(assignments, dbAssignmentToDomainAssignment(assignment))
	}

	return assignments, nil
}

func (s *Store) EditAssignment(ctx context.Context, assignment domain.Assignment) error {
	err := s.Queries.EditAssignment(ctx, db.EditAssignmentParams{
		AssignedUserID: assignment.AssignedUserID,
		ScheduledFor:   assignment.ScheduledFor,
		UpdatedAt:      assignment.UpdatedAt,
		ID:             assignment.ID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) CancelAssignment(ctx context.Context, assignment domain.Assignment) error {
	err := s.Queries.CancelAssignment(ctx, db.CancelAssignmentParams{
		Canceled:   assignment.Canceled,
		CanceledAt: sql.NullTime{Valid: true, Time: assignment.CanceledAt},
		UpdatedAt:  assignment.UpdatedAt,
		ID:         assignment.ID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) CompleteAssignment(ctx context.Context, assignment domain.Assignment) error {
	err := s.Queries.CompleteAssignment(ctx, db.CompleteAssignmentParams{
		Completed:   assignment.Completed,
		CompletedAt: sql.NullTime{Valid: true, Time: assignment.CompletedAt},
		UpdatedAt:   assignment.UpdatedAt,
		ID:          assignment.ID,
	})

	if err != nil {
		return err
	}

	return nil
}
