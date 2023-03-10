package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	transactionController "POS-System/controllers/transaction"
	userController "POS-System/controllers/user"
)

type ControllerList struct {
	UserController        userController.UserController
	TransactionController transactionController.TransactionController
	JWTMiddleware         middleware.JWTConfig
}

func (cl *ControllerList) RouteRegister(echo *echo.Echo) {
	// user
	user := echo.Group("api/v1/user")
	user.POST("/login", cl.UserController.LoginUser)

	// transactions
	transactions := echo.Group("api/v1/transactions")
	transactions.Use(middleware.JWTWithConfig(cl.JWTMiddleware))
	transactions.POST("", cl.TransactionController.CreateNewTransaction)
	transactions.GET("", cl.TransactionController.GetAllTransactions)
	transactions.DELETE("/:id", cl.TransactionController.DeleteTransaction)
	transactions.PUT("/:id", cl.TransactionController.UpdateTransaction)
}
