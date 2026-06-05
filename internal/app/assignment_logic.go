package app

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mikelawson03/chores/internal/domain"
)

// Compare requesting user to user assignment and ensure they match (return error if not)
// Make and call helper to build Assignment add App.Assignments

func (a *App) validateAssignmentRequest(choreID string, assignedUserID string, scheduleDate *time.Time) error {
	if strings.TrimSpace(choreID) == "" {
		return errors.New("Missing required chore ID")
	}
	if strings.TrimSpace(assignedUserID) == "" {
		return errors.New("Missing required assigned user ID")
	}
	if scheduleDate == nil {
		return errors.New("Missing required schedule date")
	}

	if scheduleDate.Before(time.Now()) {
		return errors.New("Chores may not be scheduled before current time")
	}
	return nil
}

func (a *App) AddAssignment(choreID string, assignedUserID string, scheduleDate *time.Time) domain.Assignment {
	id := uuid.NewString()
	assignment := domain.Assignment{
		ID:             id,
		ChoreID:        choreID,
		AssignedUserID: assignedUserID,
		ScheduledFor:   *scheduleDate,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	a.Assignments = append(a.Assignments, assignment)
	return assignment
}

func (a *App) CreateAssignmentForUser(requesterID string, choreID string, assignedUserID string, scheduleDate *time.Time) (domain.Assignment, error) {
	// validate request body
	err := a.validateAssignmentRequest(choreID, assignedUserID, scheduleDate)
	if err != nil {
		return domain.Assignment{}, err
	}

	// check authorization
	if requesterID != assignedUserID {
		return domain.Assignment{}, fmt.Errorf("Requester %s may not assign chores to user %s", requesterID, assignedUserID)
	}

	// check that chore template exists
	if !a.choreExists(choreID) {
		return domain.Assignment{}, fmt.Errorf("Unknown chore ID: %s", choreID)
	}

	assignment := a.AddAssignment(choreID, assignedUserID, scheduleDate)

	return assignment, nil
}
