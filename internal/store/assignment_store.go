package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store/db"
)

func dbAssignmentToDomainAssignment(dbAssignment db.Assignment) domain.Assignment {
	var completedAt *time.Time
	var canceledAt *time.Time

	if dbAssignment.CompletedAt.Valid {
		t := dbAssignment.CompletedAt.Time
		completedAt = &t
	}

	if dbAssignment.CanceledAt.Valid {
		t := dbAssignment.CanceledAt.Time
		completedAt = &t
	}

	return domain.Assignment{
		ID:             dbAssignment.ID,
		TemplateID:     dbAssignment.TemplateID,
		AssignedUserID: dbAssignment.AssignedUserID,
		DueDate:        dbAssignment.DueDate,
		Completed:      dbAssignment.Completed,
		Canceled:       dbAssignment.Canceled,
		CreatedAt:      dbAssignment.CreatedAt,
		UpdatedAt:      dbAssignment.UpdatedAt,
		CompletedAt:    completedAt,
		CanceledAt:     canceledAt,
	}
}

func mapGetAssignmentRow(r db.GetAssignmentRow) domain.Assignment {
	var completedAt *time.Time
	var canceledAt *time.Time
	var scheduledFor *time.Time
	var instructions string
	var notes string

	if r.CompletedAt.Valid {
		t := r.CompletedAt.Time
		completedAt = &t
	}

	if r.CanceledAt.Valid {
		t := r.CanceledAt.Time
		canceledAt = &t
	}

	if r.Instructions.Valid {
		instructions = r.Instructions.String
	}

	if r.Notes.Valid {
		notes = r.Notes.String
	}

	if r.ScheduledFor.Valid {
		t := r.ScheduledFor.Time
		scheduledFor = &t
	}

	return domain.Assignment{
		ID:           r.ID,
		TemplateID:   r.TemplateID,
		TemplateName: r.Name,
		Cadence:      r.Cadence,
		Duration:     r.Duration,
		Instructions: instructions,
		Notes:        notes,
		DueDate:      r.DueDate,
		ScheduledFor: scheduledFor,
		Completed:    r.Completed,
		Canceled:     r.Canceled,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
		CompletedAt:  completedAt,
		CanceledAt:   canceledAt,
	}
}

func mapGetAllAssignmentsRow(r db.GetAllAssignmentsRow) domain.Assignment {
	var completedAt *time.Time
	var canceledAt *time.Time
	var scheduledFor *time.Time
	var instructions string
	var notes string

	if r.CompletedAt.Valid {
		t := r.CompletedAt.Time
		completedAt = &t
	}

	if r.CanceledAt.Valid {
		t := r.CanceledAt.Time
		canceledAt = &t
	}

	if r.Instructions.Valid {
		instructions = r.Instructions.String
	}

	if r.Notes.Valid {
		notes = r.Notes.String
	}

	if r.ScheduledFor.Valid {
		t := r.ScheduledFor.Time
		scheduledFor = &t
	}

	return domain.Assignment{
		ID:           r.ID,
		TemplateID:   r.TemplateID,
		TemplateName: r.Name,
		Cadence:      r.Cadence,
		Duration:     r.Duration,
		Instructions: instructions,
		Notes:        notes,
		DueDate:      r.DueDate,
		ScheduledFor: scheduledFor,
		Completed:    r.Completed,
		Canceled:     r.Canceled,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
		CompletedAt:  completedAt,
		CanceledAt:   canceledAt,
	}
}

func (s *Store) AddAssignment(ctx context.Context, assignment domain.Assignment) error {
	var scheduledFor sql.NullTime
	var instructions sql.NullString

	if assignment.ScheduledFor != nil {
		scheduledFor = sql.NullTime{
			Valid: true,
			Time:  *assignment.ScheduledFor,
		}
	} else {
		scheduledFor = sql.NullTime{
			Valid: false,
		}
	}

	if assignment.Instructions != "" {
		instructions = sql.NullString{
			Valid:  true,
			String: assignment.Instructions,
		}
	} else {
		instructions = sql.NullString{
			Valid: false,
		}
	}

	err := s.Queries.CreateAssignment(ctx, db.CreateAssignmentParams{
		ID:             assignment.ID,
		TemplateID:     assignment.TemplateID,
		AssignedUserID: assignment.AssignedUserID,
		Instructions:   instructions,
		DueDate:        assignment.DueDate,
		ScheduledFor:   scheduledFor,
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

func (s *Store) GetAssignment(ctx context.Context, id string) (domain.Assignment, error) {
	dbAssignment, err := s.Queries.GetAssignment(ctx, id)
	if err != nil {
		return domain.Assignment{}, err
	}

	assignment := mapGetAssignmentRow(dbAssignment)

	return assignment, nil

}

func (s *Store) GetAllAssignments(ctx context.Context) ([]domain.Assignment, error) {
	dbAssignments, err := s.Queries.GetAllAssignments(ctx)
	if err != nil {
		return []domain.Assignment{}, err
	}

	assignments := []domain.Assignment{}

	for _, dbAssignment := range dbAssignments {
		assignment := mapGetAllAssignmentsRow(dbAssignment)

		assignments = append(assignments, assignment)
	}

	return assignments, nil
}

func (s *Store) EditAssignment(ctx context.Context, assignment domain.Assignment) error {
	var scheduledFor sql.NullTime
	var notes sql.NullString
	var completedAt sql.NullTime
	var canceledAt sql.NullTime

	if assignment.ScheduledFor != nil {
		scheduledFor = sql.NullTime{
			Valid: true,
			Time:  *assignment.ScheduledFor,
		}
	} else {
		scheduledFor.Valid = false
	}

	if assignment.Notes != "" {
		notes = sql.NullString{
			Valid:  true,
			String: assignment.Notes,
		}
	} else {
		notes.Valid = false
	}

	if assignment.CompletedAt != nil {
		completedAt = sql.NullTime{
			Valid: true,
			Time:  *assignment.CompletedAt,
		}
	} else {
		completedAt.Valid = false
	}

	if assignment.CanceledAt != nil {
		canceledAt = sql.NullTime{
			Valid: true,
			Time:  *assignment.CanceledAt,
		}
	} else {
		canceledAt.Valid = false
	}

	err := s.Queries.EditAssignment(ctx, db.EditAssignmentParams{
		AssignedUserID: assignment.AssignedUserID,
		Notes:          notes,
		ScheduledFor:   scheduledFor,
		UpdatedAt:      assignment.UpdatedAt,
		ID:             assignment.ID,
		Completed:      assignment.Completed,
		Canceled:       assignment.Canceled,
		CompletedAt:    completedAt,
		CanceledAt:     canceledAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetAssignmentsByTemplateID(ctx context.Context, id string) ([]domain.Assignment, error) {
	dbAssignments, err := s.Queries.GetAssignmentsByTemplateID(ctx, id)
	if err != nil {
		return []domain.Assignment{}, err
	}

	var assignments []domain.Assignment
	for _, assignment := range dbAssignments {
		assignments = append(assignments, dbAssignmentToDomainAssignment(assignment))
	}

	return assignments, nil
}

func (s *Store) GetAssignmentsByUserID(ctx context.Context, id string) ([]domain.Assignment, error) {
	dbAssignments, err := s.Queries.GetAssignmentsByUserID(ctx, id)
	if err != nil {
		return []domain.Assignment{}, err
	}

	var assignments []domain.Assignment
	for _, assignment := range dbAssignments {
		assignments = append(assignments, dbAssignmentToDomainAssignment(assignment))
	}

	return assignments, nil
}

func (s *Store) GetAssignmentsForBalancing(ctx context.Context, horizonStart, horizonEnd, monthlyPlanningEnd time.Time) ([]domain.AssignmentWithMetadata, error) {
	dbAssignmentsWithMetadata, err := s.Queries.GetAssignmentsWithMetadataForDateRange(ctx, db.GetAssignmentsWithMetadataForDateRangeParams{
		DueDate:   horizonStart,
		DueDate_2: horizonEnd,
		DueDate_3: horizonStart,
		DueDate_4: monthlyPlanningEnd,
	})
	if err != nil {
		return []domain.AssignmentWithMetadata{}, err
	}

	var assignments []domain.AssignmentWithMetadata
	for _, dbAssignmentWithMetadata := range dbAssignmentsWithMetadata {
		var completedAt *time.Time
		var canceledAt *time.Time
		if dbAssignmentWithMetadata.CompletedAt.Valid {
			t := dbAssignmentWithMetadata.CompletedAt.Time
			completedAt = &t
		}

		if dbAssignmentWithMetadata.CanceledAt.Valid {
			t := dbAssignmentWithMetadata.CanceledAt.Time
			canceledAt = &t
		}

		assignment := domain.Assignment{
			ID:             dbAssignmentWithMetadata.ID,
			TemplateID:     dbAssignmentWithMetadata.TemplateID,
			AssignedUserID: dbAssignmentWithMetadata.AssignedUserID,
			DueDate:        dbAssignmentWithMetadata.DueDate,
			Completed:      dbAssignmentWithMetadata.Completed,
			Canceled:       dbAssignmentWithMetadata.Canceled,
			CreatedAt:      dbAssignmentWithMetadata.CreatedAt,
			UpdatedAt:      dbAssignmentWithMetadata.UpdatedAt,
			CompletedAt:    completedAt,
			CanceledAt:     canceledAt,
		}

		assignmentWithDuration := domain.AssignmentWithMetadata{
			Assignment:      assignment,
			DurationMinutes: int(dbAssignmentWithMetadata.Duration),
			Cadence:         domain.Cadence(dbAssignmentWithMetadata.Cadence),
		}

		assignments = append(assignments, assignmentWithDuration)
	}

	return assignments, nil

}

func (s *Store) BulkUpdateAssignments(ctx context.Context, assignments []domain.Assignment) error {
	tx, err := s.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	qtx := s.Queries.WithTx(tx)

	for _, assignment := range assignments {
		err := qtx.EditAssignment(ctx, db.EditAssignmentParams{
			AssignedUserID: assignment.AssignedUserID,
			CreatedAt:      assignment.CreatedAt,
			UpdatedAt:      time.Now(),
			ID:             assignment.ID,
		})
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
