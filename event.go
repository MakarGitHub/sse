package sse

import "sync"

/* Event represents an SSE (Server-Sent Events) event. */
type Event struct {
	mu   sync.Mutex
	ID   string
	Data string
}

/* SetID sets the ID of the event.
It is safe for concurrent use by multiple goroutines.*/
func (e *Event) SetID(id string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.ID = id
}

/* SetData sets the data of the event.
It is safe for concurrent use by multiple goroutines.*/
func (e *Event) SetData(data string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Data = data
}

/* NewEvent creates a new Event with the given ID, event type, and data.
The returned Event is safe for concurrent use by multiple goroutines.*/
func NewEvent(data string) *Event {
	return &Event{
		Data: data,
		mu:   sync.Mutex{},
	}
}
