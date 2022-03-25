package cqrs

import (
	"github.com/satioO/scheduler/scheduler/cqrs/commands"
	"github.com/satioO/scheduler/scheduler/cqrs/marshaler"
	"github.com/satioO/scheduler/scheduler/cqrs/message"
)

type AppConfig struct {
	// It allows you to use topic per command or one topic for every command. [todo - add to doc]
	GenerateCommandsTopic func(commandName string) string

	// CommandHandlers return command handlers which should be executed.
	CommandHandlers func(commandBus *commands.CommandBus) []commands.CommandHandler

	// CommandsPublisher is Publisher used to publish commands.
	CommandsPublisher message.Publisher

	CommandEventMarshaler marshaler.CommandEventMarshaler

	// CommandsSubscriber is constructor for subscribers which will subscribe for messages.
	// It will be called for every command handler.
	// It allows you to create separated customized Subscriber for every command handler.
	CommandsSubscriber commands.CommandsSubscriber
}

type App struct {
	commandsTopic         func(commandName string) string
	commandBus            *commands.CommandBus
	commandEventMarshaler marshaler.CommandEventMarshaler
}

func (f App) CommandBus() *commands.CommandBus {
	return f.commandBus
}

func (f App) CommandEventMarshaler() marshaler.CommandEventMarshaler {
	return f.commandEventMarshaler
}

func NewApp(config *AppConfig) (*App, error) {
	commandBus, err := commands.NewCommandBus(
		config.CommandsPublisher,
		config.GenerateCommandsTopic,
		config.CommandEventMarshaler,
	)

	if err != nil {
		panic(err)
	}

	app := &App{
		commandsTopic:         config.GenerateCommandsTopic,
		commandBus:            commandBus,
		commandEventMarshaler: config.CommandEventMarshaler,
	}

	commandProcessor, err := commands.NewCommandsProcessor(
		config.CommandHandlers(commandBus),
		config.CommandsSubscriber,
		config.GenerateCommandsTopic,
		config.CommandEventMarshaler,
	)

	if err != nil {
		panic(err)
	}

	if err := commandProcessor.AddHandlersToRouter(); err != nil {
		panic(err)
	}

	return app, nil
}
