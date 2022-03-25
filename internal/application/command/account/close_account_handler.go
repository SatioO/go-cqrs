package command

import (
	"context"

	"github.com/satioO/scheduler/scheduler/cqrs/commands"
	"github.com/sirupsen/logrus"
)

type CloseAccountHandler struct {
	commandBus *commands.CommandBus
}

func NewCloseAccountHandler() *CloseAccountHandler {
	return &CloseAccountHandler{}
}

func (h CloseAccountHandler) Handle(ctx context.Context, cmd any) error {
	logrus.Printf("Handler::: %v", cmd)
	return nil
}

func (b CloseAccountHandler) HandlerName() string {
	// this name is passed to EventsSubscriberConstructor and used to generate queue name
	return "CloseAccountHandler"
}

// NewCommand returns type of command which this handle should handle. It must be a pointer.
func (b CloseAccountHandler) NewCommand() any {
	return &CloseAccountHandler{}
}
