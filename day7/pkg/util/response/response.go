package response

import (
	"net/http"

	"github.com/harisfi/alterra-agmc/day7/internal/dto"
	"github.com/labstack/echo/v4"
)

func ErrorResponse(c echo.Context, code int) error {
	return c.JSON(code, &dto.BaseResponse[interface{}]{
		Status:  "Error",
		Message: http.StatusText(code),
	})
}

func SuccessResponse(c echo.Context, T any) error {
	return c.JSON(http.StatusOK, &dto.BaseResponse[any]{
		Status:  "Success",
		Message: "Request successfully proceed",
		Data:    T,
	})
}
