package main

import (
	"github.com/satioO/scheduler/scheduler/cqrs"
	"github.com/satioO/scheduler/scheduler/cqrs/command"
	"github.com/satioO/scheduler/scheduler/cqrs/marshaler"
	framework "github.com/satioO/scheduler/scheduler/internal/adapters/framework/rest"
	"github.com/satioO/scheduler/scheduler/kafka"
)

func main() {
	publisher, err := kafka.NewPublisher()

	config := cqrs.AppConfig{
		CommandsPublisher: publisher,
		CommandHandlers: func() []command.CommandHandler {
			return []command.CommandHandler{}
		},
		CommandEventMarshaler: marshaler.JSONMarshaler{},
		GenerateCommandsTopic: func(commandName string) string {
			return commandName
		},
	}

	app, err := cqrs.NewApp(&config)

	if err != nil {
		panic(err)
	}

	// DONE::: HTTP Adapter
	restAdapter := framework.NewHttpServer(app)
	framework.Run(restAdapter)

	// TODO::: GRPC Adapter
	// grpcAdapter := framework.NewGRPCServer(app)
	// framework.Run(grpcAdapter)
}
