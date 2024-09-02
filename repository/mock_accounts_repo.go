package repository

import (
	"SimpleBank/dao/model"

	"github.com/stretchr/testify/mock"
)

type MockAccountsRepository struct {
	mock.Mock
}

func (m *MockAccountsRepository) Create(account model.Account) {
	m.Called(account)
}

func (m *MockAccountsRepository) FindByCondition(filters map[string]interface{}) ([]model.Account, error) {
	args := m.Called(filters)
	return args.Get(0).([]model.Account), args.Error(1)
}

func (m *MockAccountsRepository) Update(accountUpdates map[string]interface{}, id int64) {
	m.Called(accountUpdates, id)
}

func (m *MockAccountsRepository) Delete(id int64) {
	m.Called(id)
}
