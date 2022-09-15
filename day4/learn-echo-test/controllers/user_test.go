package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/harisfi/alterra-agmc/day4/learn-echo-test/configs"
)

func InitEcho() *echo.Echo {
	configs.InitDB()
	e := echo.New()

	return e
}

func TestGetUsersController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "berhasil",
			path:                 "/users",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success\",\"users\":[",
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		// Assertions
		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}
