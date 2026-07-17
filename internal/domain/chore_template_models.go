package domain

import "time"

type Cadence string

const (
	CadenceDaily   Cadence = "daily"
	CadenceWeekly  Cadence = "weekly"
	CadenceMonthly Cadence = "monthly"
)

type ChoreTemplate struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Cadence   Cadence   `json:"cadence"`
	Assignee  string    `json:"assignee"`
	Duration  int       `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
