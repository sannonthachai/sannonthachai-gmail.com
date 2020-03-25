package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sannonthachai/test/services"
)

type Controller struct {
	con *services.Service
	e   *echo.Echo
}

func NewController(service *services.Service, e *echo.Echo) *Controller {
	return &Controller{
		con: service,
		e:   e,
	}
}

func (con *Controller) getItem() {
	res := con.con.GetItem()
	con.e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, res)
	})
}

func (con *Controller) InitController() {
	con.getItem()
}
