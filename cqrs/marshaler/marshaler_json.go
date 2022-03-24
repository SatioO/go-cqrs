package marshaler

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/satioO/scheduler/scheduler/cqrs/message"
)

type JSONMarshaler struct {
	NewUUID      func() string
	GenerateName func(v any) string
}

func (m JSONMarshaler) Marshal(v any) (*message.Message, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	msg := message.NewMessage(
		m.newUUID(),
		b,
	)

	return msg, nil
}

func (JSONMarshaler) Unmarshal(msg *message.Message, v any) (err error) {
	return json.Unmarshal(msg.Payload, v)
}

func (m JSONMarshaler) newUUID() string {
	if m.NewUUID != nil {
		return m.NewUUID()
	}

	// default
	return uuid.New().String()
}

func (m JSONMarshaler) Name(cmdOrEvent any) string {
	if m.GenerateName != nil {
		return m.GenerateName(cmdOrEvent)
	}

	return FullyQualifiedStructName(cmdOrEvent)
}
