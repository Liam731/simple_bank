package controller

import (
	"SimpleBank/dao/model"
	"SimpleBank/helper"
	"SimpleBank/request"
	"SimpleBank/response"
	"SimpleBank/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountsController struct {
	accountsService service.AccountsService
}

func NewAccountsController(service service.AccountsService) *AccountsController {
	return &AccountsController{accountsService: service}
}

// CreateAccounts		godoc
// @Summary			Create accounts
// @Description		Save accounts data in Db.
// @Param			accounts body request.CreateAccountsRequest true "Create accounts"
// @Produce			application/json
// @Accounts			accounts
// @Success			200 {object} response.Response{}
// @Router			/accounts [post]
func (controller *AccountsController) Create(ctx *gin.Context) {
	fmt.Println("create accounts")
	createAccountsRequest := request.CreateAccountsRequest{}
	err := ctx.ShouldBindJSON(&createAccountsRequest)

	fmt.Printf("createAccountsRequest: %+v\n", createAccountsRequest)
	helper.ErrorPanic(err)

	controller.accountsService.Create(createAccountsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   createAccountsRequest,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateAccounts		godoc
// @Summary			Update accounts
// @Description		Update accounts data.
// @Param			accountId path string true "update accounts by id"
// @Param			accounts body request.CreateAccountsRequest true  "Update accounts"
// @Accounts			accounts
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/accounts/{accountId} [patch]
func (controller *AccountsController) Update(ctx *gin.Context) {
	fmt.Println("update accounts")
	updateAccountsRequest := model.Account{}
	err := ctx.ShouldBindJSON(&updateAccountsRequest)
	helper.ErrorPanic(err)

	accountId := ctx.Param("accountId")
	id, err := strconv.Atoi(accountId)
	helper.ErrorPanic(err)

	controller.accountsService.Update(updateAccountsRequest, int64(id))

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteAccounts		godoc
// @Summary			Delete accounts
// @Description		Remove accounts data by id.
// @Produce			application/json
// @Accounts			accounts
// @Success			200 {object} response.Response{}
// @Router			/accounts/{accountID} [delete]
func (controller *AccountsController) Delete(ctx *gin.Context) {
	fmt.Println("delete accounts")
	accountId := ctx.Param("accountId")
	id, err := strconv.Atoi(accountId)
	helper.ErrorPanic(err)
	controller.accountsService.Delete(int64(id))

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdAccounts 		godoc
// @Summary				Get Single accounts by id.
// @Param				accountId path string true "update accounts by id"
// @Description			Return the tahs whoes accountId valu mathes id.
// @Produce				application/json
// @Accounts				accounts
// @Success				200 {object} response.Response{}
// @Router				/accounts/{accountId} [get]
func (controller *AccountsController) FindByCondition(ctx *gin.Context) {
	fmt.Println("findbycondition accounts")
	findAccountsRequest := model.Account{}
	accountId := ctx.Param("accountId")
	id, err := strconv.Atoi(accountId)
	helper.ErrorPanic(err)
	findAccountsRequest.ID = int64(id)
	accountResponse := controller.accountsService.FindByCondition(findAccountsRequest)
	fmt.Println("accountResponse = ", accountResponse)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   accountResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
