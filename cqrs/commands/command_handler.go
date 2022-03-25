package commands

import "context"

type CommandHandler interface {
	Handle(ctx context.Context, cmd any) error
	HandlerName() string
	NewCommand() any
}
