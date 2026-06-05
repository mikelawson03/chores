package domain

import "time"

type Assignment struct {
	ID             string
	ChoreID        string
	AssignedUserID string
	ScheduledFor   time.Time
	Completed      bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
