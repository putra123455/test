package route

import (
	"net/http"
	"os"

	"echo_golang/controllers"
	"echo_golang/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Routers() *echo.Echo {
	e := echo.New()
	godotenv.Load()
	dbHost := os.Getenv("SECRET_KEY")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte(dbHost),
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Please, login here !, localhost:8000/login")
	})

	e.POST("/login", controllers.LoginUserController)
	e.GET("/users", controllers.GetUserController, echojwt.WithConfig(config))
	e.GET("users/auth", controllers.GetUserController, echojwt.WithConfig(config))
	e.POST("/users", controllers.CreateUserController, echojwt.WithConfig(config))
	e.DELETE("/users/:id", controllers.DeleteUserController, echojwt.WithConfig(config))
	e.PUT("/users/:id", controllers.UpdateUserController, echojwt.WithConfig(config))

	return e
}
