package response

type AccountsResponse struct {
	Id       int    `json:"id"`
	Owner    string `json:"Owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}
