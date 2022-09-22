package book

import (
	"net/http"
	"strconv"

	"github.com/harisfi/alterra-agmc/day7/internal/dto"
	"github.com/harisfi/alterra-agmc/day7/internal/factory"
	r "github.com/harisfi/alterra-agmc/day7/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{service: NewService(f)}
}

func (h *handler) FindBook(c echo.Context) error {
	books, err := h.service.FindBook(c.Request().Context())
	if err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, books)
}

func (h *handler) FindBookById(c echo.Context) error {
	paramId := c.Param("id")
	ID, err := strconv.Atoi(paramId)
	if err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}

	book, err := h.service.FindBookByID(c.Request().Context(), uint(ID))
	if err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, book)
}

func (h *handler) CreateBook(c echo.Context) error {
	payload := new(dto.CreateBookRequest)
	if err := c.Bind(payload); err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}
	if err := c.Validate(payload); err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}

	book, err := h.service.CreateBook(c.Request().Context(), payload)
	if err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, book)
}

func (h *handler) UpdateBook(c echo.Context) error {
	payload := new(dto.UpdateBookRequest)
	if err := c.Bind(payload); err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}
	if err := c.Validate(payload); err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}

	paramId := c.Param("id")
	ID, err := strconv.Atoi(paramId)
	if err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}

	book, err := h.service.UpdateBook(c.Request().Context(), uint(ID), payload)
	if err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, book)
}

func (h *handler) DeleteBook(c echo.Context) error {
	paramId := c.Param("id")
	ID, err := strconv.Atoi(paramId)
	if err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}

	if err := h.service.DeleteBook(c.Request().Context(), uint(ID)); err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, nil)
}
