package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mikelawson03/chores/internal/domain"
)

func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()

	return ay == by && am == bm && ad == bd
}

func assignmentExistsForTemplateAndDate(templateID string, date time.Time, assignments []domain.Assignment) bool {
	for _, assignment := range assignments {
		if assignment.TemplateID != templateID {
			continue
		}
		if sameDay(assignment.ScheduledFor, date) {
			return true
		}
	}

	return false
}

func (a *App) dailyScheduler(horizonStart, horizonEnd time.Time, templates []domain.ChoreTemplate, existingAssignments []domain.Assignment) ([]domain.Assignment, error) {
	var newAssignments []domain.Assignment
	for _, template := range templates {
		if template.Cadence != domain.CadenceDaily {
			continue
		}
		currentDay := horizonStart
		for currentDay.Before(horizonEnd) {
			if !assignmentExistsForTemplateAndDate(template.ID, currentDay, existingAssignments) {
				assignment := a.createNewAssignment(template.ID, template.Assignee, &currentDay)
				newAssignments = append(newAssignments, assignment)
			}
			currentDay = currentDay.AddDate(0, 0, 1)
		}
	}
	return newAssignments, nil
}

func (a *App) RunScheduler(ctx context.Context) error {
	// calculate current scheduling horizon
	now := time.Now()
	current_day := now.Weekday()
	daysSinceWeekStart := (int(current_day) + 6) % 7
	startOfWeek := now.AddDate(0, 0, -daysSinceWeekStart)
	horizonStart := time.Date(
		startOfWeek.Year(),
		startOfWeek.Month(),
		startOfWeek.Day(),
		0, 0, 0, 0,
		startOfWeek.Location(),
	)
	horizonEnd := horizonStart.AddDate(0, 0, 28)

	// retrieve users
	users, err := a.Store.GetAllUsers(ctx)
	if err != nil {
		return fmt.Errorf("Error retrieving users - %s", err)
	}
	if len(users) == 0 {
		return errors.New("No users available for scheduling.")
	}

	// retrieve current templates
	tmps, err := a.Store.GetChoreTemplates(ctx)
	if err != nil {
		return fmt.Errorf("Error retrieving templates - %s", err)
	}
	if len(tmps) == 0 {
		return errors.New("No chore templates available for scheduling.")
	}

	// retrieve existing assignments
	existingAssignments, err := a.Store.GetAssignmentsByDateRange(ctx, horizonStart, horizonEnd)
	if err != nil {
		return fmt.Errorf("Error retrieving assignments - %s")
	}

	// run schedulers by cadence
	newDailyAssignments, err := a.dailyScheduler(horizonStart, horizonEnd, tmps, existingAssignments)

	// persist assignments in DB
	for _, assignment := range newDailyAssignments {
		err = a.Store.AddAssignment(ctx, assignment)
		if err != nil {
			return fmt.Errorf("Error committing new assignment to DB: %s", err)
		}
	}

	fmt.Printf("Assignments created: %d\n", len(newDailyAssignments))
	return nil
	// existingAssignments, err := a.GetAllAssignments(ctx)
	// if err != nil {
	// 	return err
	// }

}
