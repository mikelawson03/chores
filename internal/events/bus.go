package events

import "github.com/mikelawson03/chores/internal/domain"

type Bus struct {
	events chan domain.DomainEvent
}

func NewBus() *Bus {
	return &Bus{
		events: make(chan domain.DomainEvent),
	}
}
