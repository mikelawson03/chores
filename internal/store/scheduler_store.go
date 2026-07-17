package store

import (
	"context"
	"time"

	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store/db"
)

func (s *Store) GetAssignmentsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]domain.Assignment, error) {
	dbAssignments, err := s.Queries.GetAssignmentsByDateRange(ctx, db.GetAssignmentsByDateRangeParams{
		DueDate:   startDate,
		DueDate_2: endDate,
	})
	if err != nil {
		return []domain.Assignment{}, err
	}

	var assignments []domain.Assignment
	for _, assignment := range dbAssignments {
		assignments = append(assignments, dbAssignmentToDomainAssignment(assignment))
	}

	return assignments, nil
}
