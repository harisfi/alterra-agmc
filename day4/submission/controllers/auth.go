package controllers

import (
	"net/http"

	"github.com/harisfi/alterra-agmc/day4/submission/libs/database"
	"github.com/harisfi/alterra-agmc/day4/submission/models"
	"github.com/labstack/echo/v4"
)

var jToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjU5MzkyMjMsInVzZXJJZCI6MX0.I_qMVlgghBUW-tr7yqVumMMZDOSmjlLui2gYa2tASzw"

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
