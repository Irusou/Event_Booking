package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func New() *Event {
	return &Event{}
}

func (e Event) Save() {
	// later save to db
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}

func GetEvent(id int) Event {
	event := Event{}
	for _, e := range events {
		if e.ID == id {
			event = e
		}
	}
	return event
}
