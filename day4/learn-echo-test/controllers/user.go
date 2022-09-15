package controllers

import (
	"net/http"

	"github.com/harisfi/alterra-agmc/day4/learn-echo-test/libs/database"
	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, e := database.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
