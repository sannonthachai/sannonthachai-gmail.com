package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RestrictController struct {
	e *echo.Echo
}

func NewRestrictController(e *echo.Echo) *RestrictController {
	return &RestrictController{
		e: e,
	}
}

func (con *RestrictController) restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func (con *RestrictController) Initial() {
	r := con.e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", con.restricted)
}
