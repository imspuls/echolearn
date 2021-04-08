package routes

import (
	"echolearn/controllers"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

var JwtAuthentication = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(os.Getenv("JWT_SECRET")),
})

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func Init() *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, mofo!")
	})

	e.GET("user", controllers.GetAllUser, JwtAuthentication)
	e.POST("user", controllers.CreateUser)
	e.PUT("user/:id", controllers.UpdateUser)
	e.DELETE("user/:id", controllers.DeleteUser)

	e.POST("login", controllers.Login)

	return e
}
