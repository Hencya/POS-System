package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	ConfigJWT "POS-System/app/configs/auth"
	configDB "POS-System/app/configs/databases"
	_middleware "POS-System/app/middlewares/logger"
	"POS-System/businesses/transactionEntity"
	transactionController "POS-System/controllers/transaction"
	_domainFactory "POS-System/repository"

	"POS-System/app/routes"
	"POS-System/businesses/userEntity"
	userController "POS-System/controllers/user"
)

func main() {
	var (
		db  = configDB.SetupDatabaseConnection()
		jwt = ConfigJWT.SetupJwt()
	)
	timeoutDur, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	timeoutContext := time.Duration(timeoutDur) * time.Millisecond

	echoApp := echo.New()

	//middleware
	echoApp.Use(middleware.CORS())
	echoApp.Use(middleware.LoggerWithConfig(_middleware.LoggerConfig()))

	//user
	userRepo := _domainFactory.NewUserRepository(db)
	userService := userEntity.NewUserServices(userRepo, &jwt, timeoutContext)
	userCtrl := userController.NewUserController(userService, &jwt)

	//transaction
	transactionRepo := _domainFactory.NewTransactionRepository(db)
	transactionService := transactionEntity.NewTransactionServices(transactionRepo, &jwt, timeoutContext)
	transactionCtrl := transactionController.NewTransactionController(transactionService, &jwt)

	//routes
	routesInit := routes.ControllerList{
		JWTMiddleware:         jwt.Init(),
		UserController:        *userCtrl,
		TransactionController: *transactionCtrl,
	}
	routesInit.RouteRegister(echoApp)

	port := os.Getenv("PORT")
	log.Fatal(echoApp.Start(":" + port))
}
