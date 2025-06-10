package main

import (
	"boilerplate-feature/app/config"
	"boilerplate-feature/app/database"
	"boilerplate-feature/app/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(cfg)
	database.InitialMigration(db)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(db, e, cfg)
	e.Logger.Fatal(e.Start(":80"))

}
