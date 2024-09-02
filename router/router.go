package router

import (
	"SimpleBank/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(accountsController *controller.AccountsController) *gin.Engine {
	router := gin.Default()
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	accountsRouter := baseRouter.Group("/accounts")
	accountsRouter.GET("/:accountId", accountsController.FindByCondition)
	accountsRouter.PUT("", accountsController.Create)
	accountsRouter.PATCH("/:accountId", accountsController.Update)
	accountsRouter.DELETE("/:accountId", accountsController.Delete)

	return router
}
