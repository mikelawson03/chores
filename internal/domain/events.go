package domain

import (
	"time"
)

type DomainEvent interface {
	EventType() string
	GetHouseholdID() string
	GetPayload() any
}

type Event[T EventPayload] struct {
	Type        string
	OccurredAt  time.Time
	HouseholdID string
	Actor       Actor
	Payload     T
}

type EventPayload interface {
	Assignment | ChoreTemplate | User
}

type AssignmentEventType string

const (
	AssignmentCreated     AssignmentEventType = "assignment.created"
	AssignmentEdited      AssignmentEventType = "assignment.edited"
	AssignmentCompleted   AssignmentEventType = "assignment.completed"
	AssignmentRescheduled AssignmentEventType = "assignment.rescheduled"
)

type Actor struct {
	Type ActorType
	ID   string
}

type ActorType string

const (
	ActorTypeUser   ActorType = "user"
	ActorTypeSystem ActorType = "system"
)

func (e Event[T]) EventType() string {
	return e.Type
}

func (e Event[T]) GetHouseholdID() string {
	return e.HouseholdID
}

func (e Event[T]) GetPayload() any {
	return e.Payload
}
