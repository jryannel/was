package was

import (
	"time"
)

type Message[T any] struct {
	// The time that the message was sent.
	Time time.Duration
	// Subject of the message.
	Subject string
	// The message type.
	Type string
	// The message data.
	Data T
}
