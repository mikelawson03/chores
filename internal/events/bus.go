package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/mikelawson03/chores/internal/auth"
	"github.com/mikelawson03/chores/internal/domain"
)

const subscriberBufferSize = 50

type Bus struct {
	events      chan domain.DomainEvent
	subscribers map[string]Subscriber
	subscribe   chan SubscribeRequest
	unsubscribe chan string
}

type Subscriber struct {
	SubscriberId  string
	HouseholdUser domain.HouseholdUser
	HouseholdId   string
	Type          SubscriberType
	Events        chan domain.DomainEvent
}

type SubscriberType string

const (
	SubscriberTypeClient SubscriberType = "client"
	SubscriberTypeLogger SubscriberType = "logger"
)

type SubscribeRequest struct {
	Subscriber Subscriber
	Done       chan struct{}
}

func NewBus() *Bus {
	return &Bus{
		events:      make(chan domain.DomainEvent),
		subscribers: make(map[string]Subscriber),
		subscribe:   make(chan SubscribeRequest),
		unsubscribe: make(chan string),
	}
}

func (b *Bus) Publish(event domain.DomainEvent) {
	b.events <- event
}

func (b *Bus) Listen() {
	timeout, _ := time.ParseDuration("5s")

	for {
		select {
		case event := <-b.events:
			for id, subscriber := range b.subscribers {
				if !shouldDeliver(event, subscriber) {
					continue
				}

				select {
				case subscriber.Events <- event:

				case <-time.After(timeout):
					b.removeSubscriber(id)
				}

			}
		case req := <-b.subscribe:
			b.subscribers[req.Subscriber.SubscriberId] = req.Subscriber
			close(req.Done)
		case subscriberID := <-b.unsubscribe:
			b.removeSubscriber(subscriberID)
		}

	}

}

func (b *Bus) Subscribe(user domain.HouseholdUser, subType SubscriberType) (string, <-chan domain.DomainEvent) {
	subscriberId := uuid.NewString()

	s := Subscriber{
		SubscriberId:  subscriberId,
		HouseholdId:   user.HouseholdID,
		HouseholdUser: user,
		Type:          subType,
		Events:        make(chan domain.DomainEvent, subscriberBufferSize),
	}

	done := make(chan struct{})

	req := SubscribeRequest{
		Subscriber: s,
		Done:       done,
	}

	b.subscribe <- req

	<-done

	return subscriberId, s.Events

}

func (b *Bus) Unsubscribe(subscriberId string) {
	b.unsubscribe <- subscriberId
}

func (b *Bus) removeSubscriber(subscriberId string) {
	subscriber, ok := b.subscribers[subscriberId]

	if !ok {
		return
	}

	close(subscriber.Events)
	delete(b.subscribers, subscriberId)
}

func shouldDeliver(event domain.DomainEvent, sub Subscriber) bool {

	if sub.Type == SubscriberTypeLogger {
		return true
	}

	if sub.Type == SubscriberTypeClient && event.GetHouseholdID() == sub.HouseholdId {
		switch payload := event.GetPayload().(type) {
		case domain.Assignment:
			return auth.CanGetAssignment(sub.HouseholdUser, payload)
		case domain.ChoreTemplate:
			return auth.CanGetChoreTemplate(sub.HouseholdUser)
		default:
			// TODO: Log unexpected/unsupported event payload type.
			return false
		}

	}

	return false
}
