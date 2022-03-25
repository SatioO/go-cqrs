package kafka

import (
	"context"

	"github.com/satioO/scheduler/scheduler/cqrs/message"
	"github.com/sirupsen/logrus"
)

type Subscriber struct{}

func NewSubscriber() (*Subscriber, error) {
	return &Subscriber{}, nil
}

func (*Subscriber) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	logrus.Printf("Topic:::%s", topic)
	return nil, nil
}

// Close should flush unsent messages, if Subscriber is async.
func (*Subscriber) Close() error {
	return nil
}
