package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	r "github.com/harisfi/alterra-agmc/day10/pkg/util/response"
	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	var jwtKey = os.Getenv("JWT_KEY")

	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return r.ErrorResponse(c, http.StatusUnauthorized)
		}

		splitToken := strings.Split(authToken, "Bearer ")
		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}

			return []byte(jwtKey), nil
		})

		if !token.Valid || err != nil {
			return r.ErrorResponse(c, http.StatusUnauthorized)
		}

		if Authorized := token.Claims.(jwt.MapClaims)["authorized"]; Authorized != nil {
			if authorized := Authorized.(bool); !authorized {
				return r.ErrorResponse(c, http.StatusUnauthorized)
			}
		} else {
			if err != nil {
				return r.ErrorResponse(c, http.StatusUnauthorized)
			}
		}

		return next(c)
	}
}
