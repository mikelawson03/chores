package domain

import "time"

type User struct {
	ID        string
	Username  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
