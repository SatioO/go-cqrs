package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/satioO/scheduler/scheduler/cqrs/message"
)

type Marshaler interface {
	Marshal(topic string, msg *message.Message) (*sarama.ProducerMessage, error)
}

// Unmarshaler unmarshals Kafka's message to Watermill's message.
type Unmarshaler interface {
	Unmarshal(*sarama.ConsumerMessage) (*message.Message, error)
}

type MarshalerUnmarshaler interface {
	Marshaler
	Unmarshaler
}

const UUIDHeaderKey = "_watermill_message_uuid"

type DefaultMarshaler struct{}

func (DefaultMarshaler) Marshal(topic string, msg *message.Message) (*sarama.ProducerMessage, error) {

	headers := []sarama.RecordHeader{{
		Key:   []byte(UUIDHeaderKey),
		Value: []byte(msg.UUID),
	}}

	return &sarama.ProducerMessage{
		Topic:   topic,
		Value:   sarama.ByteEncoder(msg.Payload),
		Headers: headers,
	}, nil
}

func (DefaultMarshaler) Unmarshal(kafkaMsg *sarama.ConsumerMessage) (*message.Message, error) {
	var messageID string

	msg := message.NewMessage(messageID, kafkaMsg.Value)

	return msg, nil
}
