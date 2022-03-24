package main

import (
	"github.com/satioO/scheduler/scheduler/cqrs"
	"github.com/satioO/scheduler/scheduler/cqrs/command"
	framework "github.com/satioO/scheduler/scheduler/internal/adapters/framework/rest"
)

func main() {
	config := cqrs.AppConfig{
		CommandHandlers: func() []command.CommandHandler {
			return nil
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
