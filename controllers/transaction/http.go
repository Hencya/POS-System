package transaction

import (
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

	"POS-System/app/middlewares/auth"
	"POS-System/businesses/transactionEntity"
	"POS-System/controllers/transaction/request"
	"POS-System/controllers/transaction/response"
	"POS-System/helpers"
)

type TransactionController struct {
	transactionService transactionEntity.Service
	jwtAuth            *auth.ConfigJWT
}

func NewTransactionController(s transactionEntity.Service, jwtAuth *auth.ConfigJWT) *TransactionController {
	return &TransactionController{
		transactionService: s,
		jwtAuth:            jwtAuth,
	}
}

func (ctrl *TransactionController) CreateNewTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Transaction{}
	res := response.Transaction{}

	user := auth.GetUser(c)
	Username := user.Username

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("The Data You Entered is Wrong", http.StatusBadRequest,
				err, helpers.EmptyObj{}))
	}

	domain := transactionEntity.Domain{}
	copier.Copy(&domain, &req)
	domain.CreatedBy = Username
	domain.Date = time.Now()

	data, err := ctrl.transactionService.CreateNewTransaction(ctx, &domain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator", http.StatusInternalServerError,
				err, helpers.EmptyObj{}))
	}
	copier.Copy(&res, &data)
	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created Recipe", http.StatusCreated,
			res))
}
