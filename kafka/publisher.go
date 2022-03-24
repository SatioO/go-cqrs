package kafka

import (
	"github.com/satioO/scheduler/scheduler/cqrs/message"
	"github.com/sirupsen/logrus"
)

type Publisher struct{}

func NewPublisher() (*Publisher, error) {
	return &Publisher{}, nil
}

func (*Publisher) Publish(topic string, messages ...*message.Message) error {
	logrus.Printf("Topic:::%s, Message: %v", topic, messages)
	return nil
}

// Close should flush unsent messages, if publisher is async.
func (*Publisher) Close() error {
	return nil
}
