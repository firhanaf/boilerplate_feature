package router

import (
	"boilerplate-feature/app/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c echo.Echo, config *config.AppConfig) {
	Userdata := userData.New(db)
	Userservice := userService.New(&Userdata)
	UserhandlerAPI := userHandler.New(&Userservice)

}
