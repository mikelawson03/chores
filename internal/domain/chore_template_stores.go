package domain

type ChoreTemplate struct {
	ID       string
	Name     string
	Cadence  string
	Shared   bool
	Assignee string
	Duration int
}
