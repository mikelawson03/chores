package store

import (
	"context"
	"time"

	"github.com/mikelawson03/chores/internal/store/db"
)

type ExistingAssignment struct {
	ID         string
	TemplateID string
	DueDate    time.Time
}

func (s *Store) GetAssignmentsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]ExistingAssignment, error) {
	dbAssignments, err := s.Queries.GetAssignmentsByDateRange(ctx, db.GetAssignmentsByDateRangeParams{
		DueDate:   startDate,
		DueDate_2: endDate,
	})
	if err != nil {
		return []ExistingAssignment{}, err
	}

	var assignments []ExistingAssignment
	for _, assignment := range dbAssignments {
		assignments = append(assignments, ExistingAssignment{
			ID:         assignment.ID,
			TemplateID: assignment.TemplateID,
			DueDate:    assignment.DueDate,
		})
	}

	return assignments, nil
}
