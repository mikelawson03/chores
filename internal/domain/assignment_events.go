package domain

import "time"

type AssignmentEventType string

const (
	AssignmentCreated         AssignmentEventType = "assignment.created"
	AssignmentEdited          AssignmentEventType = "assignment.edited"
	AssignmentCompleted       AssignmentEventType = "assignment.completed"
	AssignmentRescheduled     AssignmentEventType = "assignment.rescheduled"
	AssignmentRefreshRequired AssignmentEventType = "assignment.refreshRequired"
)

type AssignmentRefreshRequiredPayload struct {
	Reason string `json:"reason"`
}

type AssignmentRefreshReason string

const (
	AssignmentRefreshTemplateUpdated AssignmentRefreshReason = "template_updated"
)

func NewAssignmentEvent(
	eventType AssignmentEventType,
	occurredAt time.Time,
	actorType ActorType,
	assignment Assignment,
	householdID,
	userID string) Event[Assignment] {
	e := Event[Assignment]{
		Type:        string(eventType),
		OccurredAt:  occurredAt,
		HouseholdID: householdID,
		Actor: Actor{
			Type: actorType,
			ID:   userID,
		},
		Payload: assignment,
	}

	return e
}

func NewAssignmentsRefreshRequiredEvent(
	householdID string,
	actorType ActorType,
	reason AssignmentRefreshReason,
	occurredAt time.Time,
	reqUserID string) Event[AssignmentRefreshRequiredPayload] {
	return Event[AssignmentRefreshRequiredPayload]{
		Type:        string(AssignmentRefreshRequired),
		OccurredAt:  occurredAt,
		HouseholdID: householdID,
		Actor: Actor{
			Type: actorType,
			ID:   reqUserID,
		},
		Payload: AssignmentRefreshRequiredPayload{
			Reason: string(reason),
		},
	}
}
