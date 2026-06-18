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
	// Get all assignments
	assignments, err := a.Store.GetAssignmentsWithMetadata(ctx)
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

	populateCurrentLoads(userLoads, assignments)

	// Get lists of assignments by cadence
	dailyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceDaily)
	newDailyAssignments := balanceUserLoads(dailyAssignments, userLoads, domain.CadenceDaily)
	weeklyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceWeekly)
	newWeeklyAssignments := balanceUserLoads(weeklyAssignments, userLoads, domain.CadenceWeekly)
	monthlyAssignments := sortedAssignmentsForCadence(assignments, domain.CadenceMonthly)
	newMonthlyAssignments := balanceUserLoads(monthlyAssignments, userLoads, domain.CadenceMonthly)

	// Get
	for _, assignment := range newDailyAssignments {
		fmt.Println(assignment)
	}
	for _, assignment := range newWeeklyAssignments {
		fmt.Println(assignment)
	}
	for _, assignment := range newMonthlyAssignments {
		fmt.Println(assignment)
	}
	for _, user := range userLoads {
		fmt.Println(user.userId)
		fmt.Println("Daily Load: ", user.dailyLoad)
		fmt.Println("Weekly Load: ", user.weeklyLoad)
		fmt.Println("Monthly Load: ", user.monthlyLoad)
	}

	return nil
}
