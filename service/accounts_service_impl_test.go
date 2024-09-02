package service

import (
	"SimpleBank/dao/model"
	"SimpleBank/repository"
	"SimpleBank/request"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestAccountsServiceImpl_Create(t *testing.T) {
	mockRepo := new(repository.MockAccountsRepository)
	validate := validator.New()

	accountService := NewAccountsServiceImpl(mockRepo, validate)

	// 建立模擬輸入
	createRequest := request.CreateAccountsRequest{
		Owner:    "John Doe",
		Balance:  1000,
		Currency: "USD",
	}

	// 預期的 model.Account
	expectedAccount := model.Account{
		Owner:    createRequest.Owner,
		Balance:  createRequest.Balance,
		Currency: createRequest.Currency,
	}

	// 設置 mock 行為
	mockRepo.On("Create", expectedAccount).Return()

	// 呼叫 Service 的 Create 方法
	accountService.Create(createRequest)

	// 確認 mock 方法是否按預期被調用
	mockRepo.AssertExpectations(t)
}

func TestAccountsServiceImpl_FindByCondition(t *testing.T) {
	mockRepo := new(repository.MockAccountsRepository)
	validate := validator.New()

	accountService := NewAccountsServiceImpl(mockRepo, validate)

	// 模擬輸入
	searchCriteria := model.Account{
		ID:       1,
		Owner:    "John Doe",
		Balance:  1000,
		Currency: "USD",
	}

	// 預期的 filter map
	expectedFilters := map[string]interface{}{
		"id":       searchCriteria.ID,
		"owner":    searchCriteria.Owner,
		"balance":  searchCriteria.Balance,
		"currency": searchCriteria.Currency,
	}

	// 模擬 Repository 回傳的資料
	expectedAccounts := []model.Account{
		{
			ID:       1,
			Owner:    "John Doe",
			Balance:  1000,
			Currency: "USD",
		},
	}

	// 設置 mock 行為
	mockRepo.On("FindByCondition", expectedFilters).Return(expectedAccounts, nil)

	// 呼叫 Service 的 FindByCondition 方法
	accounts := accountService.FindByCondition(searchCriteria)

	// 驗證回傳值
	assert.Equal(t, expectedAccounts, accounts)

	// 確認 mock 方法是否按預期被調用
	mockRepo.AssertExpectations(t)
}
