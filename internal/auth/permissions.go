package auth

import (
	"github.com/mikelawson03/chores/internal/domain"
)

func CanAssignToUser(user domain.HouseholdUser, assignedUserID string) bool {
	if user.Role != domain.RoleAdmin && user.User.ID != assignedUserID {
		return false
	}

	return true
}

func CanGetUserAssignments(user domain.HouseholdUser, userID string) bool {
	if user.Role != domain.RoleAdmin && user.User.ID != userID {
		return false
	}

	return true
}

func CanGetAssignment(user domain.HouseholdUser, assignment domain.Assignment) bool {
	if user.Role != domain.RoleAdmin && user.User.ID != assignment.AssignedUserID {
		return false
	}

	return true
}

func CanEditAssignment(user domain.HouseholdUser, assignment domain.Assignment) bool {
	if user.Role != domain.RoleAdmin && user.User.ID != assignment.AssignedUserID {
		return false
	}

	return true
}

func CanGetUser(reqUser domain.HouseholdUser, userID string) bool {
	if reqUser.Role != domain.RoleAdmin && reqUser.User.ID != userID {
		return false
	}
	return true
}

func CanEditUser(reqUser domain.User, userID string) bool {
	return reqUser.ID == userID
}

func CanEditRole(reqUser domain.HouseholdUser) bool {
	if reqUser.Role != domain.RoleAdmin {
		return false
	}
	return true
}

func CanEditHouseholdUser(reqUser domain.HouseholdUser, userID string) bool {
	if reqUser.Role != domain.RoleAdmin && reqUser.User.ID != userID {
		return false
	}
	return true
}

func CanEditIsActive(reqUser domain.HouseholdUser) bool {
	if reqUser.Role != domain.RoleAdmin {
		return false
	}
	return true
}
