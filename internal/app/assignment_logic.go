package app

import (
	"errors"
	"fmt"
	"slices"
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

func (a *App) findAssignmentIndex(id string) int {
	for i, assignment := range a.Assignments {
		if assignment.ID == id {
			return i
		}
	}
	return -1
}

func (a *App) AddAssignment(choreID string, assignedUserID string, scheduleDate *time.Time) domain.Assignment {
	id := uuid.NewString()
	assignment := domain.Assignment{
		ID:             id,
		TemplateID:     choreID,
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

	// check authorization (may only assign to self)
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

func (a *App) GetAllAssignments() ([]domain.Assignment, error) {
	return a.Assignments, nil
}

func (a *App) GetAssignmentByID(id string) (domain.Assignment, error) {
	idx := a.findAssignmentIndex(id)
	if idx == -1 {
		return domain.Assignment{}, fmt.Errorf("Assignment with ID %s not found.", id)
	}

	return a.Assignments[idx], nil
}

func (a *App) EditAssignment(requesterID string, assignmentID string, choreID string, assignedUserID string, scheduleDate *time.Time) (domain.Assignment, error) {
	idx := a.findAssignmentIndex(assignmentID)
	if idx == -1 {
		return domain.Assignment{}, fmt.Errorf("Assignment with ID %s not found.", assignmentID)
	}

	if a.Assignments[idx].AssignedUserID != requesterID {
		return domain.Assignment{}, fmt.Errorf("Requester %s may not edit assigned chores for user %s.", requesterID, a.Assignments[idx].AssignedUserID)
	}

	err := a.validateAssignmentRequest(choreID, assignedUserID, scheduleDate)
	if err != nil {
		return domain.Assignment{}, err
	}

	assignment := domain.Assignment{
		ID:             assignmentID,
		TemplateID:     choreID,
		AssignedUserID: assignedUserID,
		ScheduledFor:   *scheduleDate,
		Completed:      a.Assignments[idx].Completed,
		CreatedAt:      a.Assignments[idx].CreatedAt,
		UpdatedAt:      time.Now(),
	}

	a.Assignments[idx] = assignment

	return assignment, nil
}

func (a *App) DeleteAssignment(assignmentID string, requesterID string) error {
	idx := a.findAssignmentIndex(assignmentID)

	if idx == -1 {
		return fmt.Errorf("Assignment with ID %s not found.", assignmentID)
	}

	if a.Assignments[idx].AssignedUserID != requesterID {
		return fmt.Errorf("Requester %s may not edit assigned chores for user %s.", requesterID, a.Assignments[idx].AssignedUserID)
	}

	a.Assignments = slices.Delete(a.Assignments, idx, idx+1)

	return nil
}

func (a *App) CompleteAssignment(assignmentID string, requesterID string) (domain.Assignment, error) {
	idx := a.findAssignmentIndex(assignmentID)

	if idx == -1 {
		return domain.Assignment{}, fmt.Errorf("Assignment with ID %s not found.", assignmentID)
	}

	if a.Assignments[idx].AssignedUserID != requesterID {
		return domain.Assignment{}, fmt.Errorf("Requester %s may not edit assigned chores for user %s.", requesterID, a.Assignments[idx].AssignedUserID)
	}

	a.Assignments[idx].Completed = true
	a.Assignments[idx].CompletedAt = time.Now()

	return a.Assignments[idx], nil
}
