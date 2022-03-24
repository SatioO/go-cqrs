package cqrs

import "github.com/satioO/scheduler/scheduler/cqrs/command"

type AppConfig struct {
	CommandHandlers func() []command.CommandHandler
}

type App struct {
	commandBus *CommandBus
}

func (f App) CommandBus() *CommandBus {
	return f.commandBus
}

func NewApp(config *AppConfig) (*App, error) {
	commandBus, err := NewCommandBus()

	if err != nil {
		panic(err)
	}

	app := &App{commandBus}

	return app, nil
}
