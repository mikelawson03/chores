package domain

import "time"

type Role string

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"firstName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"udpatedAt"`
}

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type HouseholdUser struct {
	HouseholdID string    `json:"householdId"`
	Role        Role      `json:"role"`
	DisplayName string    `json:"displayName"`
	ColorOption int       `json:"colorOption"`
	JoinedAt    time.Time `json:"joinedAt"`
	IsActive    bool      `json:"isActive"`
	User        User      `json:"user"`
}

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleUser:
		return true
	default:
		return false
	}
}
