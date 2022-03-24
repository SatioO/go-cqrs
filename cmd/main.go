package main

import (
	"github.com/satioO/scheduler/scheduler/cqrs"
	"github.com/satioO/scheduler/scheduler/cqrs/command"
	"github.com/satioO/scheduler/scheduler/cqrs/marshaler"
	framework "github.com/satioO/scheduler/scheduler/internal/adapters/framework/rest"
)

func main() {
	config := cqrs.AppConfig{
		CommandHandlers: func() []command.CommandHandler {
			return []command.CommandHandler{
				
			}
		},
		CommandEventMarshaler: marshaler.JSONMarshaler{},
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
