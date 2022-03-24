package message

import "context"

type Payload []byte

type Message struct {
	UUID    string
	Payload Payload
	ctx     context.Context
}

func NewMessage(uuid string, payload Payload) *Message {
	return &Message{
		UUID:    uuid,
		Payload: payload,
	}
}

func (m *Message) Context() context.Context {
	if m.ctx != nil {
		return m.ctx
	}
	return context.Background()
}

// SetContext sets provided context to the message.
func (m *Message) SetContext(ctx context.Context) {
	m.ctx = ctx
}
