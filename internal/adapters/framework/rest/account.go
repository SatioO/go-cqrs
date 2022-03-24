package framework

import (
	"context"
	"net/http"

	command "github.com/satioO/scheduler/scheduler/internal/application/command/account"
)

func (h *HttpServer) OpenAccount(w http.ResponseWriter, r *http.Request) {
	h.app.CommandBus().Send(context.Background(), &command.OpenAccountCommand{
		Id:             1,
		AccountHolder:  "Vaibhav Satam",
		AccountType:    "SAVINGS",
		OpeningBalance: 10000,
	})

	w.Write([]byte("Command Executed"))
}
