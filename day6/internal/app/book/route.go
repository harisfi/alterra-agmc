package book

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.FindBook)
	g.GET("/:id", h.FindBookById)
	g.POST("/:id", h.CreateBook)
	g.PUT("/:id", h.UpdateBook)
	g.DELETE("/:id", h.DeleteBook)
}
