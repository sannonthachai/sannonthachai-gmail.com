package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sannonthachai/test/controllers"
	"github.com/sannonthachai/test/database"
	"github.com/sannonthachai/test/repositories"
	"github.com/sannonthachai/test/services"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	db := database.NewDatabase()
	repo := repositories.NewRepo(db)
	service := services.NewService(repo)
	jwtService := services.NewJwtService()
	controller := controllers.NewController(e, service)
	jwtController := controllers.NewJwtController(e, jwtService)
	controller.InitController()
	jwtController.Initial()

	e.Logger.Fatal(e.Start(":1323"))
}
