package main

import (
	"context"
	"time"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/satioO/scheduler/scheduler/cqrs"
	"github.com/satioO/scheduler/scheduler/cqrs/commands"
	"github.com/satioO/scheduler/scheduler/cqrs/marshaler"
	"github.com/satioO/scheduler/scheduler/cqrs/message"
	framework "github.com/satioO/scheduler/scheduler/internal/adapters/framework/rest"
	command "github.com/satioO/scheduler/scheduler/internal/application/command/account"
	"github.com/satioO/scheduler/scheduler/kafka"
)

func main() {
	logger := watermill.NewStdLogger(false, false)

	router, err := message.NewRouter(message.RouterConfig{
		CloseTimeout: time.Minute,
	}, logger)

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"localhost:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		})

	if err != nil {
		panic(err)
	}

	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	commandsSubscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{"localhost:9092"},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         "test_consumer_group",
		},
		logger,
	)

	if err != nil {
		panic(err)
	}

	config := cqrs.AppConfig{
		GenerateCommandsTopic: func(commandName string) string {
			return commandName
		},
		CommandsPublisher: publisher,
		CommandsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			return commandsSubscriber, nil
		},
		CommandHandlers: func(cb *commands.CommandBus) []commands.CommandsHandler {
			return []commands.CommandsHandler{
				command.OpenAccountHandler{},
				command.CloseAccountHandler{},
			}
		},
		CommandEventMarshaler: marshaler.JSONMarshaler{},
		Router:                router,
		Logger:                logger,
	}

	app, err := cqrs.NewApp(&config)

	if err != nil {
		panic(err)
	}

	// processors are based on router, so they will work when router will start
	go func() {
		if err := router.Run(context.Background()); err != nil {
			panic(err)
		}
	}()

	// DONE::: HTTP Adapter
	restAdapter := framework.NewHttpServer(app)
	framework.Run(restAdapter)

	// TODO::: GRPC Adapter
	// grpcAdapter := framework.NewGRPCServer(app)
	// framework.Run(grpcAdapter)
}
