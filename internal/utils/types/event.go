package types

import "time"

type Event[T any] struct {
	Attributes map[string]string
	Data       T
	Timestamp  time.Time
}

type StreamResult[T any] struct {
	Result T
	Error  error
}
