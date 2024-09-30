package sse

import (
	"fmt"
	"net/http"
	"sync"
)

// StreamSSE is a type that handles Server-Sent Events (SSE) over HTTP.
type StreamSSE struct {
	w  http.ResponseWriter
	mu sync.Mutex
}

// Send sends a single Server-Sent Event (SSE) to the client.
//
// The function accepts an Event pointer as a parameter. If the event is nil,
// it returns an error. If the response writer's Content-Type header is not set to
// "text/event-stream", it returns an error.
//
// The function constructs the SSE message based on the event's ID and Data fields.
// It then locks the mutex to ensure thread safety while writing to the response writer.
// After writing the SSE message, it flushes the response writer to send the message to the client.
//
// If any error occurs during the process, it is returned. Otherwise, nil is returned.
func (s *StreamSSE) Send(event *Event) error {
	if s.w.Header().Get("Content-Type") != "text/event-stream" {
		return ErrResponseWriterNotSSE
	}

	if event == nil {
		return ErrNilEvent
	}

	res := ""
	if event.ID != "" {
		res = fmt.Sprintf("id: %s\n", event.ID)
	}
	if event.Data != "" {
		res += fmt.Sprintf("data: %s\n\n", event.Data)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.w.Write([]byte(res))
	if err != nil {
		return err
	}
	s.w.(http.Flusher).Flush()

	return nil
}

// NewStreamSSE creates a new instance of StreamSSE.
//
// The function accepts an http.ResponseWriter as a parameter. It sets the necessary headers
// for SSE communication and returns a new StreamSSE instance.
func NewStreamSSE(w http.ResponseWriter) *StreamSSE {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")

	return &StreamSSE{
		w:  w,
		mu: sync.Mutex{},
	}
}
