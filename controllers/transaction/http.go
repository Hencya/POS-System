package transaction

import (
	"fmt"
	"net/http"
	"strconv"

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

	data, err := ctrl.transactionService.CreateNewTransaction(ctx, &domain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator", http.StatusInternalServerError,
				err, helpers.EmptyObj{}))
	}
	copier.Copy(&res, &data)
	return c.JSON(http.StatusCreated,
		helpers.BuildSuccessResponse("Successfully created Transaction", http.StatusCreated,
			res))
}

func (ctrl *TransactionController) GetAllTransactions(c echo.Context) error {
	query := c.QueryParams()
	var page, filterAmount int
	var filterType, TypeAmount, qSort string
	offset := 0
	limit := 10
	params := transactionEntity.ParamGetTransactions{}

	if pageStr := query.Get("page"); pageStr != "" {
		page, _ = strconv.Atoi(pageStr)

		if page <= 0 {
			page = 1
		}
	}

	if offsetStr := query.Get("offset"); offsetStr != "" {
		offset, _ = strconv.Atoi(offsetStr)
	}

	if limitStr := query.Get("limit"); limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	filter := query.Get("filter")
	switch filter {
	case "min_amount":
		if minAmountStr := query.Get("min_amount"); minAmountStr != "" {
			minAmount, _ := strconv.Atoi(minAmountStr)
			filterAmount = minAmount
			TypeAmount = "min"
		}
	case "max_amount":
		if maxAmountStr := query.Get("max_amount"); maxAmountStr != "" {
			maxAmount, _ := strconv.Atoi(maxAmountStr)
			filterAmount = maxAmount
			TypeAmount = "max"
		}
	}

	filterType = query.Get("type")

	sort := query.Get("sort")
	switch sort {
	case "amount_asc":
		qSort = "amount ASC"
	case "amount_desc":
		qSort = "amount DESC"
	case "date_asc":
		qSort = "date ASC"
	case "date_desc":
		qSort = "date DESC"
	}

	params = transactionEntity.ParamGetTransactions{
		Page:       page,
		Offset:     offset,
		Limit:      limit,
		Amount:     filterAmount,
		TypeAmount: TypeAmount,
		Type:       filterType,
		Sort:       qSort,
	}
	fmt.Println("params controller", params)
	data, offsetAfterGet, totalData, err := ctrl.transactionService.GetTransactions(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Transaction Doesn't Exist", http.StatusNotFound,
				err, helpers.EmptyObj{}))
	}

	res := []response.Transaction{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offsetAfterGet,
		TotalData: totalData,
	}

	copier.Copy(&res, &data)

	if len(*data) == 0 {
		return c.JSON(http.StatusNoContent,
			helpers.BuildSuccessResponse("No transactions have been made", http.StatusOK,
				data))
	}

	return c.JSON(http.StatusOK, helpers.BuildSuccessPageResponse(http.StatusOK, "success get all transactions", res, resPage))
}

func (ctrl *TransactionController) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, errGet := ctrl.transactionService.GetTransactionById(c.Request().Context(), uint(id))
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Transaction doesn't exist", http.StatusNotFound,
				errGet, helpers.EmptyObj{}))
	}

	_, err := ctrl.transactionService.DeleteTransactionById(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator", http.StatusInternalServerError,
				err, helpers.EmptyObj{}))
	}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully Deleted a Transaction", http.StatusOK,
			nil))
}

func (ctrl *TransactionController) UpdateTransaction(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	req := request.Transaction{}
	res := response.Transaction{}

	user := auth.GetUser(c)
	Username := user.Username

	_, errGet := ctrl.transactionService.GetTransactionById(c.Request().Context(), uint(id))
	if errGet != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("Transaction doesn't exist", http.StatusNotFound,
				errGet, helpers.EmptyObj{}))
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("The Data You Entered is Wrong", http.StatusBadRequest,
				err, helpers.EmptyObj{}))
	}

	domain := transactionEntity.Domain{}
	copier.Copy(&domain, &req)
	domain.UpdatedBy = Username

	data, err := ctrl.transactionService.UpdateTransactionById(ctx, &domain, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helpers.BuildErrorResponse("Something Gone Wrong,Please Contact Administrator", http.StatusInternalServerError,
				err, helpers.EmptyObj{}))
	}
	copier.Copy(&res, &data)
	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("Successfully updated Transaction", http.StatusOK,
			res))
}
