package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/harisfi/alterra-agmc/day4/submission/libs/database"
	m "github.com/harisfi/alterra-agmc/day4/submission/middlewares"
	"github.com/harisfi/alterra-agmc/day4/submission/models"
)

func GetAllUsers(c echo.Context) error {
	users, e := database.GetAllUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all users",
		"users":   users,
	})
}

func GetUserById(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	user, e := database.GetUserById(uint(id))
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get user by id",
		"user":    user,
	})
}

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	if e := c.Validate(user); e != nil {
		return e
	}

	userCreated, e := database.CreateUser(user)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create new user",
		"user":    userCreated,
	})
}

func UpdateUserById(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	if id != m.ExtractTokenUserId(c) {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	userUpdated, e := database.UpdateUserById(uint(id), user)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update user by id",
		"user":    userUpdated,
	})
}

func DeleteUserById(c echo.Context) error {
	id, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	if id != m.ExtractTokenUserId(c) {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	if e := database.DeleteUserById(uint(id)); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete user by id",
	})
}
