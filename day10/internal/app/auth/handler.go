package auth

import (
	"net/http"

	"github.com/harisfi/alterra-agmc/day10/internal/dto"
	"github.com/harisfi/alterra-agmc/day10/internal/factory"
	r "github.com/harisfi/alterra-agmc/day10/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{service: NewService(f)}
}

func (h *handler) Login(c echo.Context) error {
	payload := new(dto.AuthLoginRequest)

	if err := c.Bind(payload); err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}
	if err := c.Validate(payload); err != nil {
		return r.ErrorResponse(c, http.StatusUnauthorized)
	}

	data, err := h.service.Login(c.Request().Context(), payload)
	if err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, data)
}

func (h *handler) Register(c echo.Context) error {
	payload := new(dto.AuthRegisterRequest)

	if err := c.Bind(payload); err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}
	if err := c.Validate(payload); err != nil {
		return r.ErrorResponse(c, http.StatusUnauthorized)
	}

	data, err := h.service.Register(c.Request().Context(), payload)
	if err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, data)
}
