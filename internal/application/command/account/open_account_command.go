package command

type OpenAccountCommand struct {
	Id             int64  `json:"id,omitempty"`
	AccountHolder  string `json:"account_holder,omitempty"`
	AccountType    string `json:"account_type,omitempty"`
	OpeningBalance int64  `json:"opening_balance,omitempty"`
}
