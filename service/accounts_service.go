package service

import (
	"SimpleBank/dao/model"
	"SimpleBank/request"
)

type AccountsService interface {
	Create(accounts request.CreateAccountsRequest)
	FindByCondition(accounts model.Account) []model.Account
	Update(accounts model.Account, id int64)
	Delete(id int64)
}
