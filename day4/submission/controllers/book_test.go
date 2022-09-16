package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/harisfi/alterra-agmc/day4/submission/configs"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/suite"
)

type BookTestSuite struct {
	suite.Suite
	Echo *echo.Echo
}

var jToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}

func (s *BookTestSuite) SetupTest() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
	}

	configs.InitDB()
	s.Echo = echo.New()
	s.Echo.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
}

func (s *BookTestSuite) TestGetAllBooks() {
	path := "/v1/books"
	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_get_all_books",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"books":[{`,
			token:                jToken,
		},
		{
			name:                 "failed_get_all_books",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"message":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.GET(path, GetAllBooks)

			req := httptest.NewRequest(http.MethodGet, path, nil)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}

func (s *BookTestSuite) TestGetBookById() {
	path := "/v1/books/1"
	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_get_book_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"book":{"id":1`,
			token:                jToken,
		},
		{
			name:                 "failed_get_book_by_id",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"message":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.GET("/v1/books/:id", GetBookById)

			req := httptest.NewRequest(http.MethodGet, path, nil)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}

func (s *BookTestSuite) TestCreateBook() {
	path := "/v1/books"
	payload := `"title":"GO","author":"John","publisher":"ABC"`

	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_create_book",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"book":{"id":3,` + payload,
			token:                jToken,
		},
		{
			name:                 "failed_create_book",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"message":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.POST(path, CreateBook)

			req := httptest.NewRequest(http.MethodPost, path, strings.NewReader("{"+payload+"}"))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}

func (s *BookTestSuite) TestUpdateBookById() {
	path := "/v1/books/1"
	payload := `"title":"GO","author":"John","publisher":"ABC"`

	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_update_book_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"book":{"id":1,` + payload,
			token:                jToken,
		},
		{
			name:                 "failed_update_book_by_id",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"message":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.PUT("/v1/books/:id", UpdateBookById)

			req := httptest.NewRequest(http.MethodPut, path, strings.NewReader("{"+payload+"}"))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}

func (s *BookTestSuite) TestDeleteBookById() {
	path := "/v1/books/1"

	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_delete_book_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"message":"Success delete book by id"}`,
			token:                jToken,
		},
		{
			name:                 "failed_delete_book_by_id",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"message":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.DELETE("/v1/books/:id", DeleteBookById)

			req := httptest.NewRequest(http.MethodDelete, path, nil)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}
