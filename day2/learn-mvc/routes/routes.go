package routes

import (
	"github.com/harisfi/alterra-agmc/day2/learn-mvc/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUserControllers)

	return e
}
