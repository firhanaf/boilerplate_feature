package handler

import (
	"boilerplate-feature/app/middlewares"
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

func (handler *UserHandler) Login(c echo.Context) error {
	var loginInput LoginRequest
	err := c.Bind(&loginInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid identifier or password", nil))
	}
	validate := validator.New()
	errVal := validate.Struct(loginInput)
	if errVal != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, errVal.Error(), nil))
	}

	token, errToken := handler.userService.Login(loginInput.Identifier, loginInput.Password)
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid identifier or password", nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success login", token))
}

func (handler *UserHandler) Me(c echo.Context) error {
	user, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, "unauthorized", nil))
	}

	result, errProfile := handler.userService.GetProfile(user)
	if errProfile != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "failed to get user profile", nil))
	}
	resultCore := UserCoretoResponse(result)
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success get user", resultCore))
}

func (handler *UserHandler) UpdateProfile(c echo.Context) error {
	// Ambil ID dari JWT Token
	userID, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, "unauthorized", nil))
	}

	var userInput UserUpdateRequest
	err = c.Bind(&userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	validate := validator.New()
	errVal := validate.Struct(userInput)
	if errVal != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "validation error", nil))
	}

	// Mapping ke UserCore
	var userCore = UserUpdateRequesttoCore(userInput)
	userCore.ID = userID // Pastikan ID dari JWT, bukan dari input user!

	errHandler := handler.userService.UpdateProfile(userID, userCore)
	if errHandler != nil {
		if strings.Contains(errHandler.Error(), "validation error") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, errHandler.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, errHandler.Error(), nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success update user", nil))
}

func (handler *UserHandler) DeleteProfile(c echo.Context) error {
	userID, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, "unauthorized", nil))
	}
	errHandler := handler.userService.DeleteAccount(userID)
	if errHandler != nil {
		if strings.Contains(errHandler.Error(), "validation error") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, errHandler.Error(), nil))
		}
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete user", nil))
}

func (handler *UserHandler) GetAllProfile(c echo.Context) error {
	userID, err := middlewares.ExtractToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusForbidden, "unauthorized", nil))
	}
	result, errHandler := handler.userService.GetAllUsers(userID)
	if errHandler != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "failed to get user profile", nil))
	}
	var userResponse []UserResponse
	for _, v := range result {
		userResponse = append(userResponse, UserCoretoResponse(v))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success get users", userResponse))

}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}
