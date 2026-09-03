package app

import (
	"context"
	"errors"
	"fmt"
	"log"
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

// creates new slice for assignments that match specified cadence and sorts them in descending order by the duration of the task in minutes (largest first)
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
func getLoadForCadence(load *userLoad, cadence domain.Cadence) (*int, error) {
	switch cadence {
	case domain.CadenceDaily:
		return &load.dailyLoad, nil
	case domain.CadenceWeekly:
		return &load.weeklyLoad, nil
	case domain.CadenceMonthly:
		return &load.monthlyLoad, nil
	default:
		return &load.dailyLoad, domain.ErrInvalidCadence
	}
}

// receives a list of userLoads and a list of assignments. adds assignment duration to assigned user's workload
// for the appropriate cadence
func populateCurrentLoads(userLoads []userLoad, assignments []store.BalancerAssignment) error {
	for _, assignment := range assignments {
		if assignment.AssignedUserID == "" {
			continue
		}
		for i := range userLoads {
			if userLoads[i].userId == assignment.AssignedUserID {
				load, err := getLoadForCadence(&userLoads[i], assignment.Cadence)
				if err != nil {
					return fmt.Errorf("assignment %s has invalid cadence %q", assignment.ID, assignment.Cadence)

				}
				*load += assignment.Duration

			}
		}
	}
	return nil
}

// receives list of userLoads and cadence; loops through them to find lowest load for specified cadence
// and returns pointer to that userLoad struct
func findLowestUserLoad(userLoads []userLoad, cadence domain.Cadence) (*userLoad, error) {
	lowest := &userLoads[0]

	for i := 1; i < len(userLoads); i++ {
		currentLoad, err := getLoadForCadence(&userLoads[i], cadence)
		if err != nil {
			return &userLoad{}, err
		}

		lowestLoad, err := getLoadForCadence(lowest, cadence)
		if err != nil {
			return &userLoad{}, err
		}
		if *currentLoad < *lowestLoad {
			lowest = &userLoads[i]
		}
	}

	return lowest, nil
}

// loops through assignments, ignores previously assigned, finds lowest load for cadence, assigns current
// assignment in loop to user with lowest load, retrieve cadence load pointer within userLoad struct, and
// increments it. Then adds assignment to newAssignments slice and returns that to caller
func balanceUserLoads(assignments []store.BalancerAssignment, userLoads []userLoad, cadence domain.Cadence) ([]store.BalancerAssignment, error) {
	var newAssignments []store.BalancerAssignment

	for _, assignment := range assignments {
		if assignment.AssignedUserID != "" {
			continue
		}

		lowestLoad, err := findLowestUserLoad(userLoads, cadence)
		if err != nil {
			return []store.BalancerAssignment{}, err
		}

		assignment.AssignedUserID = lowestLoad.userId

		load, err := getLoadForCadence(lowestLoad, cadence)
		if err != nil {
			return []store.BalancerAssignment{}, err
		}
		*load += assignment.Duration

		newAssignments = append(newAssignments, assignment)
	}
	return newAssignments, nil
}

func (a *App) RunBalancer(ctx context.Context) error {
	hhUser, err := CheckAdmin(ctx)
	if err != nil {
		return err
	}

	// Get planning horizon window and monthly planning end
	horizonStart, horizonEnd := a.getHorizonWindow()
	monthlyPlanningEnd := a.getMonthlyPlanningEnd(horizonStart, horizonEnd)
	log.Printf("Horizon Start: %v\nHorizonEnd: %v\nMonthly Planning End %v\n", horizonStart, horizonEnd, monthlyPlanningEnd)

	// Get all assignments
	assignments, err := a.Store.GetAssignmentsForBalancing(ctx, horizonStart, horizonEnd, monthlyPlanningEnd)
	if errors.Is(err, domain.ErrNotFound) {
		return fmt.Errorf("%w: no assignments available for balancing", domain.ErrInvalidRequest)
	}
	if err != nil {
		return err
	}

	// Get all users
	householdUsers, err := a.Store.GetHouseholdUsers(ctx, hhUser.HouseholdID)
	if err != nil {
		return err
	}
	if len(householdUsers) == 0 {
		return fmt.Errorf("%w: no users available for balancing", domain.ErrInvalidRequest)
	}

	// Create list of user workloads from retrieved users
	var userLoads []userLoad
	for _, householdUser := range householdUsers {
		userLoads = append(userLoads, userLoad{
			userId: householdUser.User.ID,
		})
	}

	// Calculate the workload for already assigned assignments
	err = populateCurrentLoads(userLoads, assignments)
	if err != nil {
		return err
	}

	// Get lists of assignments, balance unassigned by cadence, and append new assignments to persistence slice
	var allocationPlan []store.BalancerAssignment

	dailyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceDaily)
	newDailyAllocations, err := balanceUserLoads(dailyAssignments, userLoads, domain.CadenceDaily)
	if err != nil {
		return err
	}
	for _, assignment := range newDailyAllocations {
		allocationPlan = append(allocationPlan, assignment)
	}

	weeklyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceWeekly)
	newWeeklyAllocations, err := balanceUserLoads(weeklyAssignments, userLoads, domain.CadenceWeekly)
	if err != nil {
		return err
	}
	for _, assignment := range newWeeklyAllocations {
		allocationPlan = append(allocationPlan, assignment)
	}

	monthlyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceMonthly)
	newMonthlyAllocations, err := balanceUserLoads(monthlyAssignments, userLoads, domain.CadenceMonthly)
	if err != nil {
		return err
	}
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
