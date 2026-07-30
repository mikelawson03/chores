package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
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

func assignmentExistsForTemplateAndDate(templateID string, date time.Time, assignments []store.ExistingAssignment) bool {
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

func assignmentExistsForTemplateAndDateWindow(templateID string, windowStart, windowEnd time.Time, assignments []store.ExistingAssignment) bool {
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

func (a *App) dailyScheduler(ctx context.Context, horizonStart, horizonEnd time.Time, templates []domain.ChoreTemplate, existingAssignments []store.ExistingAssignment) error {
	created := 0
	for _, template := range templates {
		if template.Cadence != domain.CadenceDaily {
			continue
		}
		currentDay := horizonStart
		for currentDay.Before(horizonEnd) {
			if !assignmentExistsForTemplateAndDate(template.ID, currentDay, existingAssignments) {
				err := a.Store.AddAssignment(ctx, store.CreateAssignmentParams{
					ID:             uuid.NewString(),
					TemplateID:     template.ID,
					AssignedUserID: template.Assignee,
					Instructions:   template.Instructions,
					DueDate:        currentDay,
					ScheduledFor:   &currentDay,
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				})

				if err != nil {
					log.Printf("Daily scheduler aborted after creating %d assignments: %v",
						created,
						err)
					return err
				}

				created++
			}
			currentDay = currentDay.AddDate(0, 0, 1)
		}
	}
	log.Printf("Daily scheduler run complete: %d assignments created.\n", created)
	return nil
}

func (a *App) weeklyScheduler(ctx context.Context, horizonStart, horizonEnd time.Time, templates []domain.ChoreTemplate, existingAssignments []store.ExistingAssignment) error {
	created := 0
	for _, template := range templates {
		if template.Cadence != domain.CadenceWeekly {
			continue
		}

		for weekStart := horizonStart; weekStart.Before(horizonEnd); weekStart = weekStart.AddDate(0, 0, 7) {
			weekEnd := weekStart.AddDate(0, 0, 7)
			if !assignmentExistsForTemplateAndDateWindow(template.ID, weekStart, weekEnd, existingAssignments) {
				assignmentDate := weekEnd.AddDate(0, 0, -1)
				err := a.Store.AddAssignment(ctx, store.CreateAssignmentParams{
					ID:             uuid.NewString(),
					TemplateID:     template.ID,
					AssignedUserID: template.Assignee,
					Instructions:   template.Instructions,
					DueDate:        assignmentDate,
					ScheduledFor:   nil,
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				})
				if err != nil {
					log.Printf("Weekly scheduler aborted after creating %d assignments: %v",
						created,
						err)
					return err
				}
				created++
			}
		}
	}
	log.Printf("Weekly scheduler run complete: %d assignments created.\n", created)
	return nil
}

func (a *App) monthlyScheduler(ctx context.Context, horizonStart, horizonEnd time.Time, templates []domain.ChoreTemplate, existingAssignments []store.ExistingAssignment) error {
	created := 0
	for _, template := range templates {
		if template.Cadence != domain.CadenceMonthly {
			continue
		}
		thisMonthStart := time.Date(horizonStart.Year(), horizonStart.Month(), 1, 0, 0, 0, 0, horizonStart.Location())
		for thisMonthStart.Before(horizonEnd) {
			nextMonthStart := thisMonthStart.AddDate(0, 1, 0)
			thisMonthEnd := nextMonthStart.AddDate(0, 0, -1)
			if !assignmentExistsForTemplateAndDateWindow(template.ID, thisMonthStart, nextMonthStart, existingAssignments) {
				err := a.Store.AddAssignment(ctx, store.CreateAssignmentParams{
					ID:             uuid.NewString(),
					TemplateID:     template.ID,
					AssignedUserID: template.Assignee,
					Instructions:   template.Instructions,
					DueDate:        thisMonthEnd,
					ScheduledFor:   nil,
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				})
				if err != nil {
					log.Printf("Monthly scheduler aborted after creating %d assignments: %v",
						created,
						err)
					return err
				}
				created++
			}
			thisMonthStart = nextMonthStart
		}
	}
	log.Printf("Weekly scheduler run complete: %d assignments created.\n", created)
	return nil
}

func (a *App) RunScheduler(ctx context.Context) error {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return err
	}

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

	//run schedulers by cadence
	err = a.dailyScheduler(ctx, horizonStart, horizonEnd, tmps, existingAssignments)
	if err != nil {
		return err
	}
	err = a.weeklyScheduler(ctx, horizonStart, horizonEnd, tmps, existingAssignments)
	if err != nil {
		return err
	}

	err = a.monthlyScheduler(ctx, horizonStart, horizonEnd, tmps, existingAssignments)
	if err != nil {
		return err
	}

	return nil
}
