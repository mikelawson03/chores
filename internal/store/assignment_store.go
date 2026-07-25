package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store/db"
)

type CreateAssignmentParams struct {
	ID             string
	TemplateID     string
	AssignedUserID string
	Instructions   string
	DueDate        time.Time
	ScheduledFor   *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type EditAssignmentParams struct {
	ID             string
	AssignedUserID string
	ScheduledFor   *time.Time
	Notes          string
	Completed      bool
	Canceled       bool
	UpdatedAt      time.Time
	CompletedAt    *time.Time
	CanceledAt     *time.Time
}

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
		ID:                    r.ID,
		TemplateID:            r.TemplateID,
		TemplateName:          r.Name,
		AssignedUserID:        r.AssignedUserID,
		AssignedUserFirstName: r.FirstName,
		Cadence:               r.Cadence,
		Duration:              r.Duration,
		Instructions:          instructions,
		Notes:                 notes,
		DueDate:               r.DueDate,
		ScheduledFor:          scheduledFor,
		Completed:             r.Completed,
		Canceled:              r.Canceled,
		CreatedAt:             r.CreatedAt,
		UpdatedAt:             r.UpdatedAt,
		CompletedAt:           completedAt,
		CanceledAt:            canceledAt,
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

func (s *Store) AddAssignment(ctx context.Context, params CreateAssignmentParams) (domain.Assignment, error) {
	var scheduledFor sql.NullTime
	var instructions sql.NullString

	if params.ScheduledFor != nil {
		scheduledFor = sql.NullTime{
			Valid: true,
			Time:  *params.ScheduledFor,
		}
	} else {
		scheduledFor = sql.NullTime{
			Valid: false,
		}
	}

	if params.Instructions != "" {
		instructions = sql.NullString{
			Valid:  true,
			String: params.Instructions,
		}
	} else {
		instructions = sql.NullString{
			Valid: false,
		}
	}

	err := s.Queries.CreateAssignment(ctx, db.CreateAssignmentParams{
		ID:             params.ID,
		TemplateID:     params.TemplateID,
		AssignedUserID: params.AssignedUserID,
		Instructions:   instructions,
		DueDate:        params.DueDate,
		ScheduledFor:   scheduledFor,
		Completed:      false,
		Canceled:       false,
		CreatedAt:      params.CreatedAt,
		UpdatedAt:      params.UpdatedAt,
		CompletedAt:    sql.NullTime{Valid: false},
		CanceledAt:     sql.NullTime{Valid: false},
	})

	if err != nil {
		return domain.Assignment{}, err
	}

	assignment, err := s.GetAssignment(ctx, params.ID)
	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
}

func (s *Store) GetAssignment(ctx context.Context, id string) (domain.Assignment, error) {
	dbAssignment, err := s.Queries.GetAssignment(ctx, id)
	if err != nil {
		fmt.Println("hit")
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

func (s *Store) EditAssignment(ctx context.Context, editRequest EditAssignmentParams) (domain.Assignment, error) {

	err := s.Queries.EditAssignment(ctx, db.EditAssignmentParams{
		AssignedUserID: editRequest.AssignedUserID,
		Notes:          stringToNullString(editRequest.Notes),
		ScheduledFor:   pointerTimeToNullTime(editRequest.ScheduledFor),
		UpdatedAt:      editRequest.UpdatedAt,
		ID:             editRequest.ID,
		Completed:      editRequest.Completed,
		Canceled:       editRequest.Canceled,
		CompletedAt:    pointerTimeToNullTime(editRequest.CompletedAt),
		CanceledAt:     pointerTimeToNullTime(editRequest.CanceledAt),
	})

	if err != nil {
		return domain.Assignment{}, err
	}

	assignment, err := s.GetAssignment(ctx, editRequest.ID)
	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
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
			ScheduledFor:   pointerTimeToNullTime(assignment.ScheduledFor),
			Notes:          stringToNullString(assignment.Notes),
			UpdatedAt:      time.Now(),
			Completed:      assignment.Completed,
			Canceled:       assignment.Canceled,
			CompletedAt:    pointerTimeToNullTime(assignment.CompletedAt),
			CanceledAt:     pointerTimeToNullTime(assignment.CanceledAt),
			ID:             assignment.ID,
		})
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
