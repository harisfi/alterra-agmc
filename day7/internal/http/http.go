package http

import (
	"github.com/harisfi/alterra-agmc/day7/internal/app/book"
	"github.com/harisfi/alterra-agmc/day7/internal/app/user"
	"github.com/harisfi/alterra-agmc/day7/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	v1 := e.Group("/v1")
	book.NewHandler(f).Route(v1.Group("/books"))
	user.NewHandler(f).Route(v1.Group("/users"))
}
