package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sannonthachai/test/models"
	"github.com/sannonthachai/test/services"
)

type JwtController struct {
	e   *echo.Echo
	con *services.JwtService
}

func NewJwtController(e *echo.Echo, service *services.JwtService) *JwtController {
	return &JwtController{
		e:   e,
		con: service,
	}
}

func (con *JwtController) login(c echo.Context) error {
	req := new(models.JwtModel)
	c.Bind(req)
	res := con.con.LoginService(req)
	return c.JSON(http.StatusOK, res)
}

func (con *JwtController) Initial() {
	con.e.POST("/login", con.login)
}
