package domain

import "time"

type UserEventType string

const (
	UserEdited UserEventType = "user.edited"
)

func NewUserEvent(eventType UserEventType,
	occurredAt time.Time,
	actorType ActorType,
	user User,
	householdID,
	userID string) Event[User] {
	return Event[User]{
		Type:        string(eventType),
		OccurredAt:  occurredAt,
		HouseholdID: householdID,
		Actor: Actor{
			Type: actorType,
			ID:   userID,
		},
		Payload: user,
	}
}
