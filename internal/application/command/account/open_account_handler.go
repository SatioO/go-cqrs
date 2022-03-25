package command

import (
	"context"

	"github.com/satioO/scheduler/scheduler/cqrs/commands"
	"github.com/sirupsen/logrus"
)

type OpenAccountHandler struct {
	commandBus *commands.CommandBus
}

func NewOpenAccountHandler() *OpenAccountHandler {
	return &OpenAccountHandler{}
}

func (h OpenAccountHandler) Handle(ctx context.Context, cmd any) error {
	logrus.Printf("Handler::: %v", cmd)
	return nil
}

func (b OpenAccountHandler) HandlerName() string {
	// this name is passed to EventsSubscriberConstructor and used to generate queue name
	return "OpenAccountHandler"
}

// NewCommand returns type of command which this handle should handle. It must be a pointer.
func (b OpenAccountHandler) NewCommand() any {
	return &OpenAccountHandler{}
}
