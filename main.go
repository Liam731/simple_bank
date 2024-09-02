package main

import (
	"SimpleBank/config"
	"SimpleBank/controller"
	"SimpleBank/dao/model"
	"SimpleBank/helper"
	"SimpleBank/repository"
	"SimpleBank/router"
	"SimpleBank/service"
	"net/http"
	"time"

	"github.com/go-playground/validator"
)

// main is the entry point of the Go program.
//
// It sets up the database connection, initializes the repository, service, and controller,
// and starts the HTTP server.
// No parameters.
// No return values.
// main initializes the web server, sets up the router, and starts listening on port 8888.
func main() {
	// Connect to database
	db := config.DatabaseConnection()

	// Close database connection
	defer config.CloseDatabaseConnection(db)

	// Gorm CRUD
	db.AutoMigrate(&model.Account{})

	validate := validator.New()

	//Init Repository
	accountRepository := repository.NewAccountsRepositoryImpl(db)

	//Init Service
	accountService := service.NewAccountsServiceImpl(accountRepository, validate)

	//Init controller
	accountController := controller.NewAccountsController(accountService)

	//Router
	routes := router.NewRouter(accountController)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
