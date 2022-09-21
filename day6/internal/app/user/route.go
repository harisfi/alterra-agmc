package user

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.FindUser)
	g.GET("/:id", h.FindUserById)
	g.POST("/:id", h.CreateUser)
	g.PUT("/:id", h.UpdateUser)
	g.DELETE("/:id", h.DeleteUser)
}
