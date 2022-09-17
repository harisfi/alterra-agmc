package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/harisfi/alterra-agmc/day4/submission/configs"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	path := "/v1/login"
	payload := `"email":"a@b.com","password":"12345678"`

	var testCases = []struct {
		name                 string
		expectStatus         int
		payload              string
		expectBodyStartsWith string
	}{
		{
			name:                 "success_login",
			expectStatus:         http.StatusOK,
			payload:              payload,
			expectBodyStartsWith: `{"status":"Successfully logged in","token":"`,
		},
		{
			name:                 "failed_login_not_found",
			expectStatus:         http.StatusBadRequest,
			payload:              `"email":"a.b.com","password":"12345678"`,
			expectBodyStartsWith: `{"message":"record not found`,
		},
	}

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
	}

	configs.InitDB()
	configs.DB.Exec("INSERT INTO users (name,email,password) VALUES ('a','a@b.com','12345678');")

	e := echo.New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e.POST(path, Login)

			req := httptest.NewRequest(http.MethodPost, path, strings.NewReader("{"+tc.payload+"}"))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()

			e.ServeHTTP(res, req)

			assert.Equal(t, tc.expectStatus, res.Code)
			body := res.Body.String()
			assert.True(t, strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}
