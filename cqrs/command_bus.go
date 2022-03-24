package cqrs

import (
	"context"

	"github.com/sirupsen/logrus"
)

type CommandBus struct{}

func NewCommandBus() (*CommandBus, error) {
	return &CommandBus{}, nil
}

func (c CommandBus) Send(ctx context.Context, cmd any) {
	logrus.Println("Send:::")
}
