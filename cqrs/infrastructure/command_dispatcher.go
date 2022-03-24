package infrastructure

import "github.com/satioO/scheduler/scheduler/cqrs/command"

type CommandDispatcher interface {
	RegisterHandler(cmd any, handler command.CommandHandler)
	Send(cmd any)
}
