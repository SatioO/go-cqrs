package cqrs

import (
	"github.com/satioO/scheduler/scheduler/cqrs/command"
	"github.com/satioO/scheduler/scheduler/cqrs/marshaler"
)

type AppConfig struct {
	CommandHandlers       func() []command.CommandHandler
	CommandEventMarshaler marshaler.CommandEventMarshaler
}

type App struct {
	commandBus            *command.CommandBus
	commandEventMarshaler marshaler.CommandEventMarshaler
}

func (f App) CommandBus() *command.CommandBus {
	return f.commandBus
}

func (f App) CommandEventMarshaler() marshaler.CommandEventMarshaler {
	return f.commandEventMarshaler
}

func NewApp(config *AppConfig) (*App, error) {
	commandBus, err := command.NewCommandBus(config.CommandEventMarshaler)

	if err != nil {
		panic(err)
	}

	app := &App{
		commandBus:            commandBus,
		commandEventMarshaler: config.CommandEventMarshaler,
	}

	return app, nil
}
