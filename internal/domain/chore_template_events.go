package domain

import "time"

type ChoreTemplateEventType string

const (
	ChoreTemplateCreated     ChoreTemplateEventType = "choreTemplate.created"
	ChoreTemplateEdited      ChoreTemplateEventType = "choreTemplate.edited"
	ChoreTemplateDeactivated ChoreTemplateEventType = "choreTemplate.deactivated"
)

func NewChoreTemplateEvent(
	eventType ChoreTemplateEventType,
	occurredAt time.Time,
	actorType ActorType,
	choreTemplate ChoreTemplate,
	householdID,
	userID string) Event[ChoreTemplate] {
	e := Event[ChoreTemplate]{
		Type:        string(eventType),
		OccurredAt:  occurredAt,
		HouseholdID: householdID,
		Actor: Actor{
			Type: actorType,
			ID:   userID,
		},
		Payload: choreTemplate,
	}

	return e
}
