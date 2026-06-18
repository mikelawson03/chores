package app

import (
	"context"
	"fmt"
	"sort"

	"github.com/mikelawson03/chores/internal/domain"
)

type userLoad struct {
	userId      string
	dailyLoad   int
	weeklyLoad  int
	monthlyLoad int
}

// retrieve slice of assignments and cadence. creates new slice for assignments that match that cadence and
// sorts them in descending order by the duration of the task in minutes (largest first)
func sortedAssignmentsForCadence(assignments []domain.AssignmentWithMetadata, cadence domain.Cadence) []domain.AssignmentWithMetadata {
	var cadenceAssignments []domain.AssignmentWithMetadata
	for _, assignment := range assignments {
		if assignment.Cadence == cadence {
			cadenceAssignments = append(cadenceAssignments, assignment)
		}
	}
	sort.Slice(cadenceAssignments, func(i, j int) bool {
		return cadenceAssignments[i].DurationMinutes > cadenceAssignments[j].DurationMinutes
	})
	return cadenceAssignments
}

// receives pointer to userLoad struct and a cadence; returns pointer to the load value for the cadence
func getLoadForCadence(load *userLoad, cadence domain.Cadence) *int {
	switch cadence {
	case domain.CadenceDaily:
		return &load.dailyLoad
	case domain.CadenceWeekly:
		return &load.weeklyLoad
	case domain.CadenceMonthly:
		return &load.monthlyLoad
	default:
		panic(fmt.Sprintf("unknown cadence: %s", cadence))
	}
}

// receives a list of userLoads and a list of assignments. adds assignment duration to assigned user's workload
// for the appropriate cadence
func populateCurrentLoads(userLoads []userLoad, assignments []domain.AssignmentWithMetadata) {
	for _, assignment := range assignments {
		if assignment.Assignment.AssignedUserID == "" {
			continue
		}
		for i := range userLoads {
			if userLoads[i].userId == assignment.Assignment.AssignedUserID {
				load := getLoadForCadence(&userLoads[i], assignment.Cadence)
				*load += assignment.DurationMinutes

			}
		}
	}
}

// receives list of userLoads and cadence; loops through them to find lowest load for specified cadence
// and returns pointer to that userLoad struct
func findLowestUserLoad(userLoads []userLoad, cadence domain.Cadence) *userLoad {
	lowest := &userLoads[0]

	for i := 1; i < len(userLoads); i++ {
		currentLoad := getLoadForCadence(&userLoads[i], cadence)
		lowestLoad := getLoadForCadence(lowest, cadence)
		if *currentLoad < *lowestLoad {
			lowest = &userLoads[i]
		}
	}

	return lowest
}

// loops through assignments, ignores previously assigned, finds lowest load for cadence, assigns current
// assignment in loop to user with lowest load, retrieve cadence load pointer within userLoad struct, and
// increments it. Then adds assignment to newAssignments slice and returns that to caller
func balanceUserLoads(assignments []domain.AssignmentWithMetadata, userLoads []userLoad, cadence domain.Cadence) []domain.Assignment {
	var newAssignments []domain.Assignment

	for _, assignment := range assignments {
		if assignment.Assignment.AssignedUserID != "" {
			continue
		}

		lowestLoad := findLowestUserLoad(userLoads, cadence)

		assignment.Assignment.AssignedUserID = lowestLoad.userId

		load := getLoadForCadence(lowestLoad, cadence)
		*load += assignment.DurationMinutes

		newAssignments = append(newAssignments, assignment.Assignment)
	}
	return newAssignments
}

func (a *App) RunBalancer(ctx context.Context) error {
	// Get planning horizon window and monthly planning end
	horizonStart, horizonEnd := a.getHorizonWindow()
	monthlyPlanningEnd := a.getMonthlyPlanningEnd(horizonStart, horizonEnd)
	fmt.Printf("Horizon Start: %v\nHorizonEnd: %v\nMonthly Planning End %v\n", horizonStart, horizonEnd, monthlyPlanningEnd)

	// Get all assignments
	assignments, err := a.Store.GetAssignmentsForBalancing(ctx, horizonStart, horizonEnd, monthlyPlanningEnd)
	if err != nil {
		return err
	}

	// Get all users
	users, err := a.Store.GetAllUsers(ctx)
	if err != nil {
		return err
	}

	// Create list of user workloads from retrieved users
	var userLoads []userLoad
	for _, user := range users {
		userLoads = append(userLoads, userLoad{
			userId: user.ID,
		})
	}

	// Calculate the workload for already assigned assignments
	populateCurrentLoads(userLoads, assignments)

	// Get lists of assignments, balance unassigned by cadence, and append new assignments to persistence slice
	var assignmentsToPersist []domain.Assignment

	dailyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceDaily)
	newDailyAssignments := balanceUserLoads(dailyAssignments, userLoads, domain.CadenceDaily)
	for _, assignment := range newDailyAssignments {
		assignmentsToPersist = append(assignmentsToPersist, assignment)
	}

	weeklyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceWeekly)
	newWeeklyAssignments := balanceUserLoads(weeklyAssignments, userLoads, domain.CadenceWeekly)
	for _, assignment := range newWeeklyAssignments {
		assignmentsToPersist = append(assignmentsToPersist, assignment)
	}

	monthlyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceMonthly)
	newMonthlyAssignments := balanceUserLoads(monthlyAssignments, userLoads, domain.CadenceMonthly)
	for _, assignment := range newMonthlyAssignments {
		assignmentsToPersist = append(assignmentsToPersist, assignment)
	}

	// Send new assignments to persistence layer
	err = a.Store.BulkUpdateAssignments(ctx, assignmentsToPersist)
	if err != nil {
		return err
	}

	// Print metrics to console
	fmt.Println("Daily Assignments Made: ", len(newDailyAssignments))
	fmt.Println("Weekly Assignments Made: ", len(newWeeklyAssignments))
	fmt.Println("Monthly Assignments Made: ", len(newMonthlyAssignments))

	for _, user := range userLoads {
		fmt.Println(user.userId)
		fmt.Println("Daily Load: ", user.dailyLoad)
		fmt.Println("Weekly Load: ", user.weeklyLoad)
		fmt.Println("Monthly Load: ", user.monthlyLoad)
	}

	return nil
}
