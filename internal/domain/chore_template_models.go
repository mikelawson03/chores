package domain

import "time"

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
