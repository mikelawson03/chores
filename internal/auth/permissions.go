package auth

import (
	"github.com/mikelawson03/chores/internal/domain"
)

func CanAssignToUser(user domain.User, assignedUserID string) bool {
	if user.Role != domain.RoleAdmin && user.ID != assignedUserID {
		return false
	}
	return true
}

func CanGetUserAssignments(user domain.User, userID string) bool {
	if user.Role != domain.RoleAdmin && user.ID != userID {
		return false
	}

	return true
}

func CanGetAssignment(user domain.User, assignment domain.Assignment) bool {
	if user.Role != domain.RoleAdmin && user.ID != assignment.AssignedUserID {
		return false
	}

	return true
}

func CanEditAssignment(user domain.User, assignment domain.Assignment) bool {
	if user.Role != domain.RoleAdmin && user.ID != assignment.AssignedUserID {
		return false
	}

	return true
}

func CanGetUser(reqUser domain.User, userID string) bool {
	if reqUser.Role != domain.RoleAdmin && reqUser.ID != userID {
		return false
	}
	return true
}

func CanEditUserName(reqUser domain.User, userID string) bool {
	if reqUser.Role != domain.RoleAdmin && reqUser.ID != userID {
		return false
	}
	return true
}

func CanEditRole(reqUser domain.User) bool {
	if reqUser.Role != domain.RoleAdmin {
		return false
	}
	return true
}

func CanEditFirstName(reqUser domain.User, userID string) bool {
	if reqUser.Role != domain.RoleAdmin && reqUser.ID != userID {
		return false
	}
	return true
}
