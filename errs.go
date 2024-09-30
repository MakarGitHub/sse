package sse

import "fmt"

var (
	ErrResponseWriterNotSSE = fmt.Errorf("ResponseWriter is not set to SSE")
	ErrNilEvent             = fmt.Errorf("event is nil")
)
