package command

import (
	"context"
	"errors"

	"github.com/satioO/scheduler/scheduler/cqrs/marshaler"
	"github.com/sirupsen/logrus"
)

type CommandBus struct {
	marshaler marshaler.CommandEventMarshaler
}

func NewCommandBus(marshaler marshaler.CommandEventMarshaler) (*CommandBus, error) {
	if marshaler == nil {
		return nil, errors.New("missing marshaler")
	}

	return &CommandBus{
		marshaler: marshaler,
	}, nil
}

func (c CommandBus) Send(ctx context.Context, cmd any) error {
	msg, err := c.marshaler.Marshal(cmd)
	if err != nil {
		return err
	}

	commandName := c.marshaler.Name(cmd)

	msg.SetContext(ctx)

	logrus.Printf("Send::: %v", commandName)
	return nil
}
