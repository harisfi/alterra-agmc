package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    int
	Name  string
	Email string
}

func EchoMain() {
	e := echo.New()

	e.GET("/", helloController)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUser)
	e.POST("/users", createUser)
	e.POST("/users-bind", createUserBind)

	e.Start(":8080")
}

func helloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}

func getUsers(c echo.Context) error {
	users := []User{
		{1, "John", "john@doe.com"},
		{1, "John", "john@doe.com"},
	}

	match := c.QueryParam("match")
	if match != "" {
		return searchUser(c)
	}

	return c.JSON(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{id, "John", "john@doe.com"}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func searchUser(c echo.Context) error {
	match := c.QueryParam("match")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"match":  match,
		"result": []string{"adi", "aan", "asif"}, // hardcode data
	})
}

// form value
func createUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	var user User
	user.Name = name
	user.Email = email

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// binding
func createUserBind(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}
