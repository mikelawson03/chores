package domain

import "time"

type Cadence string

const (
	CadenceDaily   Cadence = "daily"
	CadenceWeekly  Cadence = "weekly"
	CadenceMonthly Cadence = "monthly"
)

type ChoreTemplate struct {
	ID        string
	Name      string
	Cadence   Cadence
	Shared    bool
	Assignee  string
	Duration  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
