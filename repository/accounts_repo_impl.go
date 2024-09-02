package repository

import (
	"SimpleBank/dao/model"
	"SimpleBank/helper"
	"errors"

	"gorm.io/gorm"
)

type AccountsRepositoryImpl struct {
	Db *gorm.DB
}

func NewAccountsRepositoryImpl(Db *gorm.DB) AccountsRepository {
	return &AccountsRepositoryImpl{Db: Db}
}
func (a AccountsRepositoryImpl) Create(accounts model.Account) {
	result := a.Db.Create(&accounts)
	helper.ErrorPanic(result.Error)
}
func (a AccountsRepositoryImpl) FindByCondition(filters map[string]interface{}) ([]model.Account, error) {
	var accounts []model.Account

	result := a.Db.Model(&model.Account{}).
		Where(filters).
		Where("created_at > ?", "2023-01-01").
		Find(&accounts)

	if result != nil {
		return accounts, nil
	} else {
		return accounts, errors.New("account is not found")
	}
}

func (a AccountsRepositoryImpl) Update(accounts map[string]interface{}, id int64) {
	result := a.Db.Model(&model.Account{}).
		Where("id = ?", id).
		Updates(accounts)

	helper.ErrorPanic(result.Error)
}

func (a AccountsRepositoryImpl) Delete(accountsId int64) {
	var accounts model.Account

	result := a.Db.Where("id = ?", accountsId).Delete(&accounts)
	helper.ErrorPanic(result.Error)
}
