package commands

import (
	"context"
	"errors"

	"github.com/satioO/scheduler/scheduler/cqrs/marshaler"
	"github.com/satioO/scheduler/scheduler/cqrs/message"
	"github.com/sirupsen/logrus"
)

type CommandsSubscriber func(handlerName string) (message.Subscriber, error)

type CommandsProcessor struct {
	handlers      []CommandHandler
	subscriber    CommandsSubscriber
	generateTopic func(commandName string) string
	marshaler     marshaler.CommandEventMarshaler
}

func NewCommandsProcessor(
	handlers []CommandHandler,
	subscriber CommandsSubscriber,
	generateTopic func(commandName string) string,
	marshaler marshaler.CommandEventMarshaler,
) (*CommandsProcessor, error) {
	if len(handlers) == 0 {
		return nil, errors.New("missing handlers")
	}

	if subscriber == nil {
		return nil, errors.New("missing subscriber")
	}

	return &CommandsProcessor{
		handlers,
		subscriber,
		generateTopic,
		marshaler,
	}, nil
}

func (p CommandsProcessor) Handlers() []CommandHandler {
	return p.handlers
}

// AddHandlersToRouter adds the CommandProcessor's handlers to the given router.
func (p CommandsProcessor) AddHandlersToRouter() error {
	for i := range p.Handlers() {
		handler := p.handlers[i]
		handlerName := handler.HandlerName()

		subscriber, err := p.subscriber(handlerName)
		if err != nil {
			return errors.New("cannot create subscriber for command processor")
		}

		messages, err := subscriber.Subscribe(context.Background(), handlerName)

		if err != nil {
			return errors.New("cannot create subscriber for command processor")
		}

		for msg := range messages {
			logrus.Print(msg)
		}
		logrus.Println(handlerName, handler.NewCommand())
	}
	return nil
}
