package service

import (
	"SimpleBank/dao/model"
	"SimpleBank/helper"
	"SimpleBank/repository"
	"SimpleBank/request"
	"fmt"

	"github.com/go-playground/validator"
)

type AccountsServiceImpl struct {
	AccountsRepository repository.AccountsRepository
	Validate           *validator.Validate
}

func NewAccountsServiceImpl(accountRepository repository.AccountsRepository, validate *validator.Validate) AccountsService {
	return &AccountsServiceImpl{
		AccountsRepository: accountRepository,
		Validate:           validate,
	}
}

func (a AccountsServiceImpl) Create(accounts request.CreateAccountsRequest) {
	err := a.Validate.Struct(accounts)
	helper.ErrorPanic(err)
	accountModel := model.Account{
		Owner:    accounts.Owner,
		Balance:  accounts.Balance,
		Currency: accounts.Currency,
	}
	a.AccountsRepository.Create(accountModel)
}

func (a AccountsServiceImpl) FindByCondition(accounts model.Account) []model.Account {
	var rtnData []model.Account
	filters := make(map[string]interface{})
	if accounts.ID != 0 {
		filters["id"] = accounts.ID
	}
	if accounts.Owner != "" {
		filters["owner"] = accounts.Owner
	}
	if accounts.Balance != 0 {
		filters["balance"] = accounts.Balance
	}
	if accounts.Currency != "" {
		filters["currency"] = accounts.Currency
	}

	accountsData, err := a.AccountsRepository.FindByCondition(filters)
	helper.ErrorPanic(err)

	for _, data := range accountsData {
		fmt.Printf("%+v\n", data)
		rtnData = append(rtnData, data)
	}

	return rtnData
}

func (a AccountsServiceImpl) Update(accounts model.Account, id int64) {
	accountUpdates := map[string]interface{}{
		"balance": accounts.Balance,
	}

	a.AccountsRepository.Update(accountUpdates, id)
}

func (a AccountsServiceImpl) Delete(id int64) {
	a.AccountsRepository.Delete(id)
}
