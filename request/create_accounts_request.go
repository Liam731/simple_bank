package request

type CreateAccountsRequest struct {
	Owner    string `json:"owner" validate:"required,min=1,max=200"`
	Currency string `json:"currency" validate:"required,min=1,max=200"`
	Balance  int64  `json:"balance" validate:"required"`
}
