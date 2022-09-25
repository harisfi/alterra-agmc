package user

import (
	"github.com/harisfi/alterra-agmc/day10/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.FindUser)
	g.POST("", h.CreateUser, middleware.Authentication)
	g.GET("/:id", h.FindUserById)
	g.PUT("/:id", h.UpdateUser, middleware.Authentication)
	g.DELETE("/:id", h.DeleteUser, middleware.Authentication)
}
