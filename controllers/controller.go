package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sannonthachai/test/models"
	"github.com/sannonthachai/test/services"
)

type Controller struct {
	e   *echo.Echo
	con *services.Service
}

func NewController(e *echo.Echo, service *services.Service) *Controller {
	return &Controller{
		con: service,
		e:   e,
	}
}

func (con *Controller) getItem(c echo.Context) error {
	res := con.con.GetItem()
	return c.JSON(http.StatusOK, res)
}

func (con *Controller) testPostModel(c echo.Context) (err error) {
	req := new(models.Model)
	c.Bind(req)
	con.con.SaveItem(req)
	return c.JSON(http.StatusCreated, req)
}

func (con *Controller) deleteItem(c echo.Context) error {
	req := new(models.DeleteModel)
	c.Bind(req)
	con.con.DeleteItem(req.Code)
	return c.JSON(http.StatusOK, "OK")
}

func (con *Controller) InitController() {
	con.e.POST("/users", con.testPostModel)
	con.e.POST("/delete_user", con.deleteItem)
	con.e.GET("/", con.getItem)
}
