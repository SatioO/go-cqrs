package commands

import (
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
func (p CommandsProcessor) AddHandlersToRouter(r *message.Router) error {
	for i := range p.Handlers() {
		handler := p.handlers[i]
		handlerName := handler.HandlerName()
		commandName := p.marshaler.Name(handler.NewCommand())
		topicName := p.generateTopic(commandName)

		subscriber, err := p.subscriber(commandName)
		if err != nil {
			return err
		}

		handlerFunc, err := p.routerHandlerFunc(handler)
		if err != nil {
			return err
		}

		r.AddNoPublisherHandler(
			handlerName,
			topicName,
			subscriber,
			handlerFunc,
		)

		logrus.Info("Router Handler", r)
	}

	return nil
}

func (p CommandsProcessor) routerHandlerFunc(handler CommandHandler) (message.NoPublishHandlerFunc, error) {
	return func(msg *message.Message) error {
		cmd := handler.NewCommand()
		messageCmdName := p.marshaler.Name(cmd)

		if err := p.marshaler.Unmarshal(msg, cmd); err != nil {
			return err
		}

		if err := handler.Handle(msg.Context(), cmd); err != nil {
			logrus.Debug("Error when handling command", err)
			return err
		}

		logrus.Printf("message_uuid: %v, received_command_type: %s", msg.UUID, messageCmdName)

		return nil
	}, nil
}
