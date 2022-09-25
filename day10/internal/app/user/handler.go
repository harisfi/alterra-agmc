package user

import (
	"net/http"
	"strconv"

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

func (h *handler) FindUser(c echo.Context) error {
	users, err := h.service.FindUser(c.Request().Context())
	if err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, users)
}

func (h *handler) FindUserById(c echo.Context) error {
	paramId := c.Param("id")
	ID, err := strconv.Atoi(paramId)
	if err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}

	user, err := h.service.FindUserByID(c.Request().Context(), uint(ID))
	if err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, user)
}

func (h *handler) CreateUser(c echo.Context) error {
	payload := new(dto.CreateUserRequest)
	if err := c.Bind(payload); err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}
	if err := c.Validate(payload); err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}

	user, err := h.service.CreateUser(c.Request().Context(), payload)
	if err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, user)
}

func (h *handler) UpdateUser(c echo.Context) error {
	payload := new(dto.UpdateUserRequest)
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

	user, err := h.service.UpdateUser(c.Request().Context(), uint(ID), payload)
	if err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, user)
}

func (h *handler) DeleteUser(c echo.Context) error {
	paramId := c.Param("id")
	ID, err := strconv.Atoi(paramId)
	if err != nil {
		return r.ErrorResponse(c, http.StatusBadRequest)
	}

	if err := h.service.DeleteUser(c.Request().Context(), uint(ID)); err != nil {
		return r.ErrorResponse(c, http.StatusInternalServerError)
	}

	return r.SuccessResponse(c, nil)
}
