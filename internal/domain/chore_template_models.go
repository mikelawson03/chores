package domain

import "time"

type ChoreTemplate struct {
	ID        string
	Name      string
	Cadence   string
	Shared    bool
	Assignee  string
	Duration  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
