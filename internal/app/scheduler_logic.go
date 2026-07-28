package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/mikelawson03/chores/internal/domain"
)

type ExistingAssignment struct {
	ID         string
	TemplateID string
	DueDate    time.Time
}

func sameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()

	return ay == by && am == bm && ad == bd
}

func schedulerAssignmentWindowEnd(horizonEnd time.Time) time.Time {
	return time.Date(
		horizonEnd.Year(),
		horizonEnd.Month(),
		1, 0, 0, 0, 0, horizonEnd.Location(),
	).AddDate(0, 1, 0)
}

func assignmentExistsForTemplateAndDate(templateID string, date time.Time, assignments []domain.Assignment) bool {
	for _, assignment := range assignments {
		if assignment.TemplateID != templateID {
			continue
		}
		if sameDay(assignment.DueDate, date) {
			return true
		}
	}

	return false
}

func assignmentExistsForTemplateAndDateWindow(templateID string, windowStart, windowEnd time.Time, assignments []domain.Assignment) bool {
	for _, assignment := range assignments {
		if assignment.TemplateID != templateID {
			continue
		}

		if assignment.DueDate.Before(windowStart) {
			continue
		}

		if assignment.DueDate.After(windowEnd) {
			continue
		}
		return true
	}
	return false
}

func (a *App) dailyScheduler(horizonStart, horizonEnd time.Time, templates []domain.ChoreTemplate, existingAssignments []domain.Assignment) []domain.Assignment {
	var newAssignments []domain.Assignment
	for _, template := range templates {
		if template.Cadence != domain.CadenceDaily {
			continue
		}
		currentDay := horizonStart
		for currentDay.Before(horizonEnd) {
			if !assignmentExistsForTemplateAndDate(template.ID, currentDay, existingAssignments) {
				assignment := a.createNewAssignment(template.ID, template.Assignee, template.Instructions, &currentDay, &currentDay)
				newAssignments = append(newAssignments, assignment)
			}
			currentDay = currentDay.AddDate(0, 0, 1)
		}
	}
	return newAssignments
}

func (a *App) weeklyScheduler(horizonStart, horizonEnd time.Time, templates []domain.ChoreTemplate, existingAssignments []domain.Assignment) []domain.Assignment {
	var newAssignments []domain.Assignment
	for _, template := range templates {
		if template.Cadence != domain.CadenceWeekly {
			continue
		}

		for weekStart := horizonStart; weekStart.Before(horizonEnd); weekStart = weekStart.AddDate(0, 0, 7) {
			weekEnd := weekStart.AddDate(0, 0, 7)
			if !assignmentExistsForTemplateAndDateWindow(template.ID, weekStart, weekEnd, existingAssignments) {
				assignmentDate := weekEnd.AddDate(0, 0, -1)
				assignment := a.createNewAssignment(template.ID, template.Assignee, template.Instructions, &assignmentDate, nil)
				newAssignments = append(newAssignments, assignment)
			}
		}
	}

	return newAssignments
}

func (a *App) monthlyScheduler(horizonStart, horizonEnd time.Time, templates []domain.ChoreTemplate, existingAssignments []domain.Assignment) []domain.Assignment {
	var newAssignments []domain.Assignment
	for _, template := range templates {
		if template.Cadence != domain.CadenceMonthly {
			continue
		}
		thisMonthStart := time.Date(horizonStart.Year(), horizonStart.Month(), 1, 0, 0, 0, 0, horizonStart.Location())
		for thisMonthStart.Before(horizonEnd) {
			nextMonthStart := thisMonthStart.AddDate(0, 1, 0)
			thisMonthEnd := nextMonthStart.AddDate(0, 0, -1)
			if !assignmentExistsForTemplateAndDateWindow(template.ID, thisMonthStart, nextMonthStart, existingAssignments) {
				assignment := a.createNewAssignment(template.ID, template.Assignee, template.Instructions, &thisMonthEnd, nil)
				newAssignments = append(newAssignments, assignment)
			}
			thisMonthStart = nextMonthStart
		}
	}

	return newAssignments
}

func (a *App) RunScheduler(ctx context.Context) error {
	// calculate current scheduling horizon
	horizonStart, horizonEnd := a.getHorizonWindow()

	// retrieve users
	users, err := a.Store.GetAllUsers(ctx)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errors.New("no users available for scheduling")
	}

	// retrieve current templates
	tmps, err := a.Store.GetChoreTemplates(ctx)
	if err != nil {
		return err
	}
	if len(tmps) == 0 {
		return errors.New("no chore templates available for scheduling")
	}

	// retrieve existing assignments
	assignmentWindowEnd := schedulerAssignmentWindowEnd(horizonEnd)
	existingAssignments, err := a.Store.GetAssignmentsByDateRange(ctx, horizonStart, assignmentWindowEnd)
	if err != nil {
		return fmt.Errorf("error retrieving assignments - %s", err)
	}

	// run schedulers by cadence
	// newDailyAssignments := a.dailyScheduler(horizonStart, horizonEnd, tmps, existingAssignments)
	// newWeeklyAssignments := a.weeklyScheduler(horizonStart, horizonEnd, tmps, existingAssignments)
	// newMonthlyAssignments := a.monthlyScheduler(horizonStart, horizonEnd, tmps, existingAssignments)

	// // // persist assignments in DB
	// // for _, dailyAssignment := range newDailyAssignments {
	// // 	_, err = a.Store.AddAssignment(ctx, dailyAssignment)
	// // 	if err != nil {
	// // 		return fmt.Errorf("error committing new daily assignment to DB: %s", err)
	// // 	}
	// // }

	// // for _, weeklyAssignment := range newWeeklyAssignments {
	// // 	_, err = a.Store.AddAssignment(ctx, weeklyAssignment)
	// // 	if err != nil {
	// // 		return fmt.Errorf("error committing new weekly assignment to DB: %s", err)
	// // 	}
	// // }

	// // for _, monthlyAssignment := range newMonthlyAssignments {
	// // 	_, err = a.Store.AddAssignment(ctx, monthlyAssignment)
	// // 	if err != nil {
	// // 		return fmt.Errorf("error committing new monthly assignment to DB: %s", err)
	// // 	}
	// // }

	// // // print new assignments created to console for debugging
	// // fmt.Printf("Daily Assignments created: %d\n", len(newDailyAssignments))
	// // fmt.Printf("Weekly Assignments created: %d\n", len(newWeeklyAssignments))
	// // fmt.Printf("Monthly Assignments created: %d\n", len(newMonthlyAssignments))
	return nil
}
