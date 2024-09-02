package repository

import (
	"SimpleBank/dao/model"
)

type AccountsRepository interface {
	Create(accounts model.Account)
	FindByCondition(filters map[string]interface{}) ([]model.Account, error)
	Update(accounts map[string]interface{}, id int64)
	Delete(accountsId int64)
}
