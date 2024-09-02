package request

type UpdateAccountsRequest struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
	Balance  int64  `json:"balance"`
}
