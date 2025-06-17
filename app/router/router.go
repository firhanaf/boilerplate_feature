package router

import (
	"boilerplate-feature/app/config"
	"boilerplate-feature/app/middlewares"
	_Userdata "boilerplate-feature/features/users/data"
	_Userhandler "boilerplate-feature/features/users/handler"
	_Userservice "boilerplate-feature/features/users/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo, config *config.AppConfig) {
	Userdata := _Userdata.New(db)
	Userservice := _Userservice.New(Userdata)
	UserhandlerAPI := _Userhandler.New(Userservice)

	api := c.Group("/api")
	api.POST("/register", UserhandlerAPI.Register)
	api.POST("/login", UserhandlerAPI.Login)
	api.GET("/me", UserhandlerAPI.Me, middlewares.JWTMiddleware())
	api.PATCH("/update", UserhandlerAPI.UpdateProfile, middlewares.JWTMiddleware())
	api.DELETE("/delete", UserhandlerAPI.DeleteProfile, middlewares.JWTMiddleware())
	api.GET("/users", UserhandlerAPI.GetAllProfile, middlewares.JWTMiddleware())
}
