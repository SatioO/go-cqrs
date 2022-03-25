package command

type CloseAccountCommand struct {
	Id             int64  `json:"id,omitempty"`
	AccountHolder  string `json:"account_holder,omitempty"`
	AccountType    string `json:"account_type,omitempty"`
	ClosingBalance int64  `json:"closing_balance,omitempty"`
}
