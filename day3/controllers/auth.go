package controllers

import (
	"net/http"

	"github.com/harisfi/alterra-agmc/day3/libs/database"
	"github.com/harisfi/alterra-agmc/day3/models"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	token, e := database.LoginUser(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Successfully logged in",
		"token":  token,
	})
}
