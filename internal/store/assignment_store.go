package store

import (
	"context"
	"database/sql"
	"errors"
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

type BalancerAssignment struct {
	ID             string
	TemplateID     string
	AssignedUserID string
	DueDate        time.Time
	Cadence        domain.Cadence
	Duration       int
}

func mapGetAssignmentsByTemplateID(r db.GetAssignmentsByTemplateIDRow) domain.Assignment {
	return domain.Assignment{
		ID:                    r.ID,
		TemplateID:            r.TemplateID,
		TemplateName:          r.Name,
		AssignedUserID:        NullStringToString(r.AssignedUserID),
		AssignedUserFirstName: NullStringToString(r.FirstName),
		Cadence:               r.Cadence,
		Duration:              r.Duration,
		Instructions:          NullStringToString(r.Instructions),
		Notes:                 NullStringToString(r.Notes),
		DueDate:               r.DueDate,
		ScheduledFor:          NullTimeToPointerTime(r.ScheduledFor),
		Completed:             r.Completed,
		Canceled:              r.Canceled,
		CreatedAt:             r.CreatedAt,
		UpdatedAt:             r.UpdatedAt,
		CompletedAt:           NullTimeToPointerTime(r.CompletedAt),
		CanceledAt:            NullTimeToPointerTime(r.CanceledAt),
	}
}

func mapGetAssignmentsByUserIDRow(r db.GetAssignmentsByUserIDRow) domain.Assignment {
	return domain.Assignment{
		ID:                    r.ID,
		TemplateID:            r.TemplateID,
		TemplateName:          r.Name,
		AssignedUserID:        NullStringToString(r.AssignedUserID),
		AssignedUserFirstName: NullStringToString(r.FirstName),
		Cadence:               r.Cadence,
		Duration:              r.Duration,
		Instructions:          NullStringToString(r.Instructions),
		Notes:                 NullStringToString(r.Notes),
		DueDate:               r.DueDate,
		ScheduledFor:          NullTimeToPointerTime(r.ScheduledFor),
		Completed:             r.Completed,
		Canceled:              r.Canceled,
		CreatedAt:             r.CreatedAt,
		UpdatedAt:             r.UpdatedAt,
		CompletedAt:           NullTimeToPointerTime(r.CompletedAt),
		CanceledAt:            NullTimeToPointerTime(r.CanceledAt),
	}
}

func mapGetAssignmentRow(r db.GetAssignmentRow) domain.Assignment {
	return domain.Assignment{
		ID:                    r.ID,
		TemplateID:            r.TemplateID,
		TemplateName:          r.Name,
		AssignedUserID:        NullStringToString(r.AssignedUserID),
		AssignedUserFirstName: NullStringToString(r.FirstName),
		Cadence:               r.Cadence,
		Duration:              r.Duration,
		Instructions:          NullStringToString(r.Instructions),
		Notes:                 NullStringToString(r.Notes),
		DueDate:               r.DueDate,
		ScheduledFor:          NullTimeToPointerTime(r.ScheduledFor),
		Completed:             r.Completed,
		Canceled:              r.Canceled,
		CreatedAt:             r.CreatedAt,
		UpdatedAt:             r.UpdatedAt,
		CompletedAt:           NullTimeToPointerTime(r.CompletedAt),
		CanceledAt:            NullTimeToPointerTime(r.CanceledAt),
	}
}

func mapGetAllAssignmentsRow(r db.GetAllAssignmentsRow) domain.Assignment {
	return domain.Assignment{
		ID:                    r.ID,
		TemplateID:            r.TemplateID,
		TemplateName:          r.Name,
		AssignedUserID:        NullStringToString(r.AssignedUserID),
		AssignedUserFirstName: NullStringToString(r.FirstName),
		Cadence:               r.Cadence,
		Duration:              r.Duration,
		Instructions:          NullStringToString(r.Instructions),
		Notes:                 NullStringToString(r.Notes),
		DueDate:               r.DueDate,
		ScheduledFor:          NullTimeToPointerTime(r.ScheduledFor),
		Completed:             r.Completed,
		Canceled:              r.Canceled,
		CreatedAt:             r.CreatedAt,
		UpdatedAt:             r.UpdatedAt,
		CompletedAt:           NullTimeToPointerTime(r.CompletedAt),
		CanceledAt:            NullTimeToPointerTime(r.CanceledAt),
	}
}

func (s *Store) AddAssignment(ctx context.Context, params CreateAssignmentParams) error {
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
		AssignedUserID: stringToNullString(params.AssignedUserID),
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
		return err
	}

	return nil
}

func (s *Store) GetAssignment(ctx context.Context, id string) (domain.Assignment, error) {
	dbAssignment, err := s.Queries.GetAssignment(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Assignment{}, fmt.Errorf("%w: assignment", domain.ErrNotFound)
	}
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

func (s *Store) EditAssignment(ctx context.Context, editRequest EditAssignmentParams) (domain.Assignment, error) {

	err := s.Queries.EditAssignment(ctx, db.EditAssignmentParams{
		AssignedUserID: stringToNullString(editRequest.AssignedUserID),
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
		assignments = append(assignments, mapGetAssignmentsByTemplateID(assignment))
	}

	return assignments, nil
}

func (s *Store) GetAssignmentsByUserID(ctx context.Context, id string) ([]domain.Assignment, error) {
	dbAssignments, err := s.Queries.GetAssignmentsByUserID(ctx, stringToNullString(id))
	if err != nil {
		return []domain.Assignment{}, err
	}

	var assignments []domain.Assignment
	for _, assignment := range dbAssignments {
		assignments = append(assignments, mapGetAssignmentsByUserIDRow(assignment))
	}

	return assignments, nil
}

func (s *Store) GetAssignmentsForBalancing(ctx context.Context,
	horizonStart,
	horizonEnd,
	monthlyPlanningEnd time.Time) ([]BalancerAssignment, error) {
	res, err := s.Queries.GetAssignmentsWithMetadataForDateRange(ctx,
		db.GetAssignmentsWithMetadataForDateRangeParams{
			DueDate:   horizonStart,
			DueDate_2: horizonEnd,
			DueDate_3: horizonStart,
			DueDate_4: monthlyPlanningEnd,
		})

	if errors.Is(err, sql.ErrNoRows) {
		return []BalancerAssignment{}, domain.ErrNotFound
	}

	if err != nil {
		return []BalancerAssignment{}, err
	}

	var assignments []BalancerAssignment
	for _, assignment := range res {

		cadence := domain.Cadence(assignment.Cadence)
		if !cadence.IsValid() {
			return []BalancerAssignment{}, domain.ErrInvalidCadence
		}

		assignment := BalancerAssignment{
			ID:             assignment.ID,
			TemplateID:     assignment.TemplateID,
			AssignedUserID: NullStringToString(assignment.AssignedUserID),
			DueDate:        assignment.DueDate,
			Cadence:        cadence,
			Duration:       int(assignment.Duration),
		}

		assignments = append(assignments, assignment)
	}

	return assignments, nil

}

func (s *Store) BulkAssignmentAllocations(ctx context.Context, assignments []BalancerAssignment) error {
	tx, err := s.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	qtx := s.Queries.WithTx(tx)

	for _, assignment := range assignments {
		err := qtx.AllocateAssignments(ctx, db.AllocateAssignmentsParams{
			AssignedUserID: stringToNullString(assignment.AssignedUserID),
			ID:             assignment.ID,
		})
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *Store) ToggleAssignmentCompletion(ctx context.Context, completion bool, completedAt, scheduledFor *time.Time, updatedAt time.Time, id string) (domain.Assignment, error) {
	err := s.Queries.UpdateAssignmentCompletion(ctx, db.UpdateAssignmentCompletionParams{
		Completed:    completion,
		CompletedAt:  pointerTimeToNullTime(completedAt),
		ScheduledFor: pointerTimeToNullTime(scheduledFor),
		UpdatedAt:    updatedAt,
		ID:           id,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Assignment{}, fmt.Errorf("%w: assignment", domain.ErrNotFound)
	}
	if err != nil {
		return domain.Assignment{}, err
	}

	assignment, err := s.GetAssignment(ctx, id)
	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
}

func (s *Store) RescheduleAssignment(ctx context.Context, assignmentID string, scheduledFor *time.Time, updatedAt time.Time) (domain.Assignment, error) {
	err := s.Queries.RescheduleAssignment(ctx, db.RescheduleAssignmentParams{
		ScheduledFor: pointerTimeToNullTime(scheduledFor),
		UpdatedAt:    updatedAt,
		ID:           assignmentID,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Assignment{}, fmt.Errorf("%w: assignment", domain.ErrNotFound)
	}
	if err != nil {
		return domain.Assignment{}, err
	}

	assignment, err := s.GetAssignment(ctx, assignmentID)
	if err != nil {
		return domain.Assignment{}, err
	}

	return assignment, nil
}

func (s *Store) GetCurrentUserAssignments(ctx context.Context, userId string) ([]domain.Assignment, error) {
	res, err := s.Queries.GetCurrentUserAssignments(ctx, stringToNullString(userId))
	if err != nil {
		return []domain.Assignment{}, err
	}

	assignments := make([]domain.Assignment, 0, len(res))
	for _, r := range res {
		assignments = append(assignments, domain.Assignment{
			ID:                    r.ID,
			TemplateID:            r.TemplateID,
			TemplateName:          r.Name,
			AssignedUserID:        NullStringToString(r.AssignedUserID),
			AssignedUserFirstName: r.FirstName,
			Cadence:               r.Cadence,
			Duration:              r.Duration,
			Instructions:          NullStringToString(r.Instructions),
			Notes:                 NullStringToString(r.Notes),
			DueDate:               r.DueDate,
			ScheduledFor:          NullTimeToPointerTime(r.ScheduledFor),
			Completed:             r.Completed,
			Canceled:              r.Canceled,
			CreatedAt:             r.CreatedAt,
			UpdatedAt:             r.UpdatedAt,
			CompletedAt:           NullTimeToPointerTime(r.CompletedAt),
			CanceledAt:            NullTimeToPointerTime(r.CanceledAt),
		})
	}

	return assignments, nil
}
