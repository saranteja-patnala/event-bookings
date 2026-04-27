package events

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time"`
	UserId      int       `json:"user_id" binding:"required"`
}

var events = []Event{}

func (e Event) SaveEvent() {
	events = append(events, e)
}

func GetEvents() []Event {
	return events
}
