package app

import (
	"context"
	"fmt"
	"sort"

	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/store"
)

type userLoad struct {
	userId      string
	dailyLoad   int
	weeklyLoad  int
	monthlyLoad int
}

// retrieve slice of assignments and cadence. creates new slice for assignments that match that cadence and
// sorts them in descending order by the duration of the task in minutes (largest first)
func sortedAssignmentsForCadence(assignments []store.BalancerAssignment, cadence domain.Cadence) []store.BalancerAssignment {
	var cadenceAssignments []store.BalancerAssignment
	for _, assignment := range assignments {
		if assignment.Cadence == cadence {
			cadenceAssignments = append(cadenceAssignments, assignment)
		}
	}
	sort.Slice(cadenceAssignments, func(i, j int) bool {
		return cadenceAssignments[i].Duration > cadenceAssignments[j].Duration
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
func populateCurrentLoads(userLoads []userLoad, assignments []store.BalancerAssignment) {
	for _, assignment := range assignments {
		if assignment.AssignedUserID == "" {
			continue
		}
		for i := range userLoads {
			if userLoads[i].userId == assignment.AssignedUserID {
				load := getLoadForCadence(&userLoads[i], assignment.Cadence)
				*load += assignment.Duration

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
func balanceUserLoads(assignments []store.BalancerAssignment, userLoads []userLoad, cadence domain.Cadence) []store.BalancerAssignment {
	var newAssignments []store.BalancerAssignment

	for _, assignment := range assignments {
		if assignment.AssignedUserID != "" {
			continue
		}

		lowestLoad := findLowestUserLoad(userLoads, cadence)

		assignment.AssignedUserID = lowestLoad.userId

		load := getLoadForCadence(lowestLoad, cadence)
		*load += assignment.Duration

		newAssignments = append(newAssignments, assignment)
	}
	return newAssignments
}

func (a *App) RunBalancer(ctx context.Context) error {
	_, err := CheckAdmin(ctx)
	if err != nil {
		return err
	}

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
	var allocationPlan []store.BalancerAssignment

	dailyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceDaily)
	newDailyAllocations := balanceUserLoads(dailyAssignments, userLoads, domain.CadenceDaily)
	for _, assignment := range newDailyAllocations {
		allocationPlan = append(allocationPlan, assignment)
	}

	weeklyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceWeekly)
	newWeeklyAllocations := balanceUserLoads(weeklyAssignments, userLoads, domain.CadenceWeekly)
	for _, assignment := range newWeeklyAllocations {
		allocationPlan = append(allocationPlan, assignment)
	}

	monthlyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceMonthly)
	newMonthlyAllocations := balanceUserLoads(monthlyAssignments, userLoads, domain.CadenceMonthly)
	for _, assignment := range newMonthlyAllocations {
		allocationPlan = append(allocationPlan, assignment)
	}

	// Send new assignments to persistence layer
	err = a.Store.BulkAssignmentAllocations(ctx, allocationPlan)
	if err != nil {
		return err
	}

	// Print metrics to console
	fmt.Println("Daily Assignments Made: ", len(newDailyAllocations))
	fmt.Println("Weekly Assignments Made: ", len(newWeeklyAllocations))
	fmt.Println("Monthly Assignments Made: ", len(newMonthlyAllocations))

	for _, user := range userLoads {
		fmt.Println(user.userId)
		fmt.Println("Daily Load: ", user.dailyLoad)
		fmt.Println("Weekly Load: ", user.weeklyLoad)
		fmt.Println("Monthly Load: ", user.monthlyLoad)
	}

	return nil
}
