package sse

import (
	"net/http/httptest"
	"testing"
)

func TestNewStreamSSE(t *testing.T) {
	w := httptest.NewRecorder()
	s := NewStreamSSE(w)

	if s == nil {
		t.Error("Nil StreamSSE returned")
	}

	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("expected Access-Control-Allow-Origin to be '*'")
	}
	if w.Header().Get("Cache-Control") != "no-cache" {
		t.Error("expected Cache-Control to be 'no-cache'")
	}
	if w.Header().Get("Connection") != "keep-alive" {
		t.Error("expected Connection to be 'keep-alive'")
	}
	if w.Header().Get("Content-Type") != "text/event-stream" {
		t.Error("expected Content-Type to be 'text/event-stream'")
	}
}
