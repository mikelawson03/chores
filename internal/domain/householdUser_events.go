package domain

import "time"

type HouseholdUserEventType string

const (
	HouseholdUserAdded          HouseholdUserEventType = "householdUser.added"
	HouseholdUserEdited         HouseholdUserEventType = "householdUser.edited"
	HouseholdUserDeactivated    HouseholdUserEventType = "householdUser.deactivated"
	HouseholdUserActivated      HouseholdUserEventType = "householdUser.activated"
	HouseholdUserRemoved        HouseholdUserEventType = "householdUser.removed"
	HouseholdUserProfileUpdated HouseholdUserEventType = "householdUser.profileUpdated"
)

type HouseholdUserProfileUpdatedPayload struct {
	UserID string `json:"userId"`
}

func NewHouseholdUserEvent(
	eventType HouseholdUserEventType,
	occurredAt time.Time,
	actorType ActorType,
	hhUser HouseholdUser,
	householdID,
	userID string) Event[HouseholdUser] {
	return Event[HouseholdUser]{
		Type:        string(eventType),
		OccurredAt:  occurredAt,
		HouseholdID: householdID,
		Actor: Actor{
			Type: actorType,
			ID:   userID,
		},
		Payload: hhUser,
	}
}

func NewHouseholdUserProfileUpdatedEvent(
	occurredAt time.Time,
	actorType ActorType,
	payload HouseholdUserProfileUpdatedPayload,
	householdID,
	reqUserID string) Event[HouseholdUserProfileUpdatedPayload] {
	return Event[HouseholdUserProfileUpdatedPayload]{
		Type:        string(HouseholdUserProfileUpdated),
		OccurredAt:  occurredAt,
		HouseholdID: householdID,
		Actor: Actor{
			Type: actorType,
			ID:   reqUserID,
		},
		Payload: payload,
	}
}
