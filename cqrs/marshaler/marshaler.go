package marshaler

import "github.com/satioO/scheduler/scheduler/cqrs/message"

type CommandEventMarshaler interface {
	// Marshal marshals Command or Event to Watermill's message.
	Marshal(v any) (*message.Message, error)

	// Unmarshal unmarshals watermill's message to v Command or Event.
	Unmarshal(msg *message.Message, v any) (err error)

	// Name returns the name of Command or Event.
	// Name is used to determine, that received command or event is event which we want to handle.
	Name(v interface{}) string
}
