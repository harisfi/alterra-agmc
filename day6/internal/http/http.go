package http

import (
	"github.com/harisfi/alterra-agmc/day6/internal/app/book"
	"github.com/harisfi/alterra-agmc/day6/internal/app/user"
	"github.com/harisfi/alterra-agmc/day6/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	book.NewHandler(f).Route(e.Group("/books"))
	user.NewHandler(f).Route(e.Group("/users"))
}
