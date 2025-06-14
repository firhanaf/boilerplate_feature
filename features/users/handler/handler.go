package handler

import (
	"boilerplate-feature/features/users"
	"boilerplate-feature/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func (handler *UserHandler) Register(c echo.Context) error {
	var userInput UserRequest
	err := c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}
	validate := validator.New()
	errVal := validate.Struct(userInput)
	if errVal != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "email/phone format is not valid", nil))
	}
	var userCore = UserRequesttoCore(userInput)
	errHandler := handler.userService.Register(userCore)
	if errHandler != nil {
		if strings.Contains(errHandler.Error(), "validation error") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, errHandler.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, errHandler.Error(), nil))
		}
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success register user", nil))

}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}
