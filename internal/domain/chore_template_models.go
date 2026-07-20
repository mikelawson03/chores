package domain

import "time"

type Cadence string

const (
	CadenceDaily   Cadence = "daily"
	CadenceWeekly  Cadence = "weekly"
	CadenceMonthly Cadence = "monthly"
)

type ChoreTemplate struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Cadence      Cadence   `json:"cadence"`
	Assignee     string    `json:"assignee"`
	Instructions string    `json:"instructions"`
	Duration     int       `json:"duration"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
