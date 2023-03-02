package userController

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"POS-System/app/middlewares/auth"
	"POS-System/businesses/userEntity"
	"POS-System/controllers/user/request"
	"POS-System/helpers"
)

type UserController struct {
	usersService userEntity.Service
	jwtAuth      *auth.ConfigJWT
}

func NewUserController(userService userEntity.Service, jwtAuth *auth.ConfigJWT) *UserController {
	return &UserController{
		usersService: userService,
		jwtAuth:      jwtAuth,
	}
}

func (ctrl *UserController) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.UserLogin{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			helpers.BuildErrorResponse("The Data You Entered is Wrong", http.StatusBadRequest,
				err, helpers.EmptyObj{}))
	}

	token, err := ctrl.usersService.Login(ctx, req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			helpers.BuildErrorResponse("User Doesn't Exist", http.StatusNotFound,
				err, helpers.EmptyObj{}))
	}

	res := struct {
		Token string `json:"token"`
	}{Token: token}

	return c.JSON(http.StatusOK,
		helpers.BuildSuccessResponse("successful to login", http.StatusOK,
			res))
}
