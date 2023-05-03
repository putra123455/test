package middleware

import (
	"echo_golang/models"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func GetClaims(c echo.Context) (*models.JwtCustomClaims, error) {
	tokenString := c.Request().Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	godotenv.Load()
	dbHost := os.Getenv("SECRET_KEY")
	token, _ := jwt.ParseWithClaims(tokenString, &models.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(dbHost), nil
	})
	claims, _ := token.Claims.(*models.JwtCustomClaims)

	return claims, nil
}
