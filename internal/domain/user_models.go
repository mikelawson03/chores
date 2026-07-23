package domain

import "time"

type Role string

type User struct {
	ID        string
	Username  string
	FirstName string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)
