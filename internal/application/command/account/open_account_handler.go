package command

type OpenAccountHandler struct{}

func NewOpenAccountHandler() *OpenAccountHandler {
	return &OpenAccountHandler{}
}

func (h *OpenAccountHandler) Handle(cmd *OpenAccountCommand) error {
	return nil
}
