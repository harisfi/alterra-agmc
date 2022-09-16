package routes

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/harisfi/alterra-agmc/day4/submission/controllers"
	m "github.com/harisfi/alterra-agmc/day4/submission/middlewares"
)

func New() *echo.Echo {
	e := echo.New()
	jwtMiddleware := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))

	v1 := e.Group("/v1")
	v1.POST("/login", controllers.Login)

	booksV1 := v1.Group("/books")
	booksV1Auth := v1.Group("/books", jwtMiddleware)

	booksV1.GET("", controllers.GetAllBooks)
	booksV1.GET("/:id", controllers.GetBookById)
	booksV1Auth.POST("", controllers.CreateBook)
	booksV1Auth.PUT("/:id", controllers.UpdateBookById)
	booksV1Auth.DELETE("/:id", controllers.DeleteBookById)

	usersV1 := v1.Group("/users")
	usersV1Auth := v1.Group("/users", jwtMiddleware)

	usersV1Auth.GET("", controllers.GetAllUsers)
	usersV1Auth.GET("/:id", controllers.GetUserById)
	usersV1.POST("", controllers.CreateUser)
	usersV1Auth.PUT("/:id", controllers.UpdateUserById)
	usersV1Auth.DELETE("/:id", controllers.DeleteUserById)

	m.LogMiddleware(e)
	e.Validator = &m.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = m.CustomErrorHandler

	return e
}
