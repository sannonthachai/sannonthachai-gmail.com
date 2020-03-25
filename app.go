package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sannonthachai/test/controllers"
	"github.com/sannonthachai/test/database"
	"github.com/sannonthachai/test/repositories"
	"github.com/sannonthachai/test/services"
)

func main() {
	e := echo.New()
	db := database.NewDatabase()
	repo := repositories.NewRepo(db)
	service := services.NewService(repo)
	controller := controllers.NewController(service, e)
	controller.InitController()

	e.Logger.Fatal(e.Start(":1323"))
}
