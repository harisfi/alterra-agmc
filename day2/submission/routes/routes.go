package routes

import (
	"github.com/harisfi/alterra-agmc/day2/submission/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")

	booksV1 := v1.Group("/books")
	booksV1.GET("", controllers.GetAllBooks)
	booksV1.GET("/:id", controllers.GetBookById)
	booksV1.POST("", controllers.CreateBook)
	booksV1.PUT("/:id", controllers.UpdateBookById)
	booksV1.DELETE("/:id", controllers.DeleteBookById)

	usersV1 := v1.Group("/users")
	usersV1.GET("", controllers.GetAllUsers)
	usersV1.GET("/:id", controllers.GetUserById)
	usersV1.POST("", controllers.CreateUser)
	usersV1.PUT("/:id", controllers.UpdateUserById)
	usersV1.DELETE("/:id", controllers.DeleteUserById)

	return e
}
