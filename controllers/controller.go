package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sannonthachai/test/models"
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
	con.e.GET("/", func(c echo.Context) error {
		res := con.con.GetItem()
		return c.JSON(http.StatusOK, res)
	})
}

func (con *Controller) testPostModel() {
	con.e.POST("/users", func(c echo.Context) (err error) {
		req := new(models.Model)
		c.Bind(req)
		con.con.SaveItem(req)
		return c.JSON(http.StatusCreated, req)
	})
}

func (con *Controller) deleteItem() {
	con.e.POST("delete_user", func(c echo.Context) error {
		req := new(models.DeleteModel)
		c.Bind(req)
		con.con.DeleteItem(req.Code)
		return c.JSON(http.StatusOK, "OK")
	})
}

func (con *Controller) InitController() {
	con.getItem()
	con.testPostModel()
	con.deleteItem()
}
