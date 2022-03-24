package command

import "context"

type CommandHandler interface {
	Handle(ctx context.Context, cmd any) error
}
