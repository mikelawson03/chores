package domain

import "time"

type Assignment struct {
	ID             string
	TemplateID     string
	AssignedUserID string
	ScheduledFor   time.Time
	Completed      bool
	Canceled       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CompletedAt    time.Time
	CanceledAt     time.Time
}
