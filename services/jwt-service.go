package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sannonthachai/test/models"
)

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (con *JwtService) LoginService(request *models.JwtModel) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString([]byte("secret"))

	return t
}
