package domain

import (
	"time"
)

type Assignment struct {
	ID                    string     `json:"id"`
	TemplateID            string     `json:"templateId"`
	TemplateName          string     `json:"templateName"`
	AssignedUserID        string     `json:"userId"`
	AssignedUserFirstName string     `json:"userFirstName"`
	Cadence               string     `json:"cadence"`
	Duration              int64      `json:"duration"`
	Instructions          string     `json:"instructions"`
	Notes                 string     `json:"notes"`
	DueDate               time.Time  `json:"dueDate"`
	ScheduledFor          *time.Time `json:"scheduledFor"`
	Completed             bool       `json:"completed"`
	Canceled              bool       `json:"canceled"`
	CreatedAt             time.Time  `json:"createdAt"`
	UpdatedAt             time.Time  `json:"updatedAt"`
	CompletedAt           *time.Time `json:"completedAt"`
	CanceledAt            *time.Time `json:"canceledAt"`
}

type AssignmentWithMetadata struct {
	Assignment      Assignment
	DurationMinutes int
	Cadence         Cadence
}
