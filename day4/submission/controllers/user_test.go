package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/harisfi/alterra-agmc/day4/submission/configs"
	"github.com/harisfi/alterra-agmc/day4/submission/middlewares"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	Echo *echo.Echo
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) SetupTest() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(err)
	}

	configs.InitDB()
	configs.DB.Exec("TRUNCATE TABLE users")

	s.Echo = echo.New()
	s.Echo.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
}

func (s *UserTestSuite) TestGetAllUsers() {
	path := "/v1/users"
	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_get_all_users",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"message":"Success get all users","users":[`,
			token:                jToken,
		},
		{
			name:                 "failed_get_all_users",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"message":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.GET(path, GetAllUsers)

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

func (s *UserTestSuite) TestGetUserById() {
	path := "/v1/users/"
	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
		withCreatedUser      bool
		userId               string
	}{
		{
			name:                 "success_get_user_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"message":"Success get user by id","user":{`,
			token:                jToken,
			withCreatedUser:      true,
			userId:               "1",
		},
		{
			name:                 "failed_get_user_by_id_not_found",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: `{"message":"record not found`,
			token:                jToken,
			userId:               "2",
		},
		{
			name:                 "failed_get_user_by_id_unauthorized",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"message":"`,
			token:                "x.x.x",
			userId:               "1",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			if tc.withCreatedUser {
				configs.DB.Exec("INSERT INTO users (name,email,password) VALUES ('a','b','c')")
			}

			s.Echo.GET(path+":id", GetUserById)

			req := httptest.NewRequest(http.MethodGet, path+tc.userId, nil)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			log.Println(body)
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}

func (s *UserTestSuite) TestCreateUser() {
	path := "/v1/users"

	var testCases = []struct {
		name                 string
		expectStatus         int
		payload              string
		expectBodyStartsWith string
		token                string
	}{
		{
			name:                 "success_create_user",
			expectStatus:         http.StatusOK,
			payload:              `"name":"John","email":"a@b.com","password":"12345678"`,
			expectBodyStartsWith: `{"message":"Success create new user","user":{"id":1,"name":"John","email":"a@b.com","password":"12345678"`,
			token:                jToken,
		},
		{
			name:                 "failed_create_user_required_name",
			expectStatus:         http.StatusBadRequest,
			payload:              `"email":"a@b.com","password":"12345678"`,
			expectBodyStartsWith: `{"message":"Name is required`,
			token:                jToken,
		},
		{
			name:                 "failed_create_user_required_email",
			expectStatus:         http.StatusBadRequest,
			payload:              `"name":"John","password":"12345678"`,
			expectBodyStartsWith: `{"message":"Email is required`,
			token:                jToken,
		},
		{
			name:                 "failed_create_user_required_password",
			expectStatus:         http.StatusBadRequest,
			payload:              `"name":"John","email":"a@b.com"`,
			expectBodyStartsWith: `{"message":"Password is required`,
			token:                jToken,
		},
		{
			name:                 "failed_create_user_false_email",
			expectStatus:         http.StatusBadRequest,
			payload:              `"name":"John","email":"a.b.com","password":"12345678"`,
			expectBodyStartsWith: `{"message":"Email is not valid emai`,
			token:                jToken,
		},
		{
			name:                 "failed_create_user_false_password",
			expectStatus:         http.StatusBadRequest,
			payload:              `"name":"John","email":"a@b.com","password":"1234"`,
			expectBodyStartsWith: `{"message":"Password value must be greater than 8`,
			token:                jToken,
		},
		{
			name:                 "failed_create_user_unauthorized",
			expectStatus:         http.StatusUnauthorized,
			expectBodyStartsWith: `{"message":"`,
			token:                "x.x.x",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.Echo.POST(path, CreateUser)
			s.Echo.Validator = &middlewares.CustomValidator{Validator: validator.New()}
			s.Echo.HTTPErrorHandler = middlewares.CustomErrorHandler

			req := httptest.NewRequest(http.MethodPost, path, strings.NewReader("{"+tc.payload+"}"))
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

func (s *UserTestSuite) TestUpdateUserById() {
	path := "/v1/users/"
	payload := `"name":"John","email":"a@b.c","password":"1234"`

	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
		withCreatedUser      bool
		userId               string
	}{
		{
			name:                 "success_update_user_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"message":"Success update user by id","user":{"id":1,` + payload,
			token:                jToken,
			withCreatedUser:      true,
			userId:               "1",
		},
		{
			name:                 "failed_update_user_by_id_forbidden",
			expectStatus:         http.StatusForbidden,
			expectBodyStartsWith: `{"message":"Forbidden`,
			token:                jToken,
			userId:               "2",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			if tc.withCreatedUser {
				configs.DB.Exec("INSERT INTO users (name,email,password) VALUES ('a','b','c')")
			}

			s.Echo.PUT(path+":id", UpdateUserById)

			req := httptest.NewRequest(http.MethodPut, path+tc.userId, strings.NewReader("{"+payload+"}"))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			log.Println(body)
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}

func (s *UserTestSuite) TestDeleteUserById() {
	path := "/v1/users/"

	var testCases = []struct {
		name                 string
		expectStatus         int
		expectBodyStartsWith string
		token                string
		withCreatedUser      bool
		userId               string
	}{
		{
			name:                 "success_delete_user_by_id",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"message":"Success delete user by id"}`,
			token:                jToken,
			withCreatedUser:      true,
			userId:               "1",
		},
		{
			name:                 "failed_delete_user_by_id_forbidden",
			expectStatus:         http.StatusForbidden,
			expectBodyStartsWith: `{"message":"Forbidden`,
			token:                jToken,
			userId:               "2",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			if tc.withCreatedUser {
				configs.DB.Exec("INSERT INTO users (name,email,password) VALUES ('a','b','c')")
			}

			s.Echo.DELETE(path+":id", DeleteUserById)

			req := httptest.NewRequest(http.MethodDelete, path+tc.userId, nil)
			req.Header.Set(echo.HeaderAuthorization, "Bearer "+tc.token)
			res := httptest.NewRecorder()

			s.Echo.ServeHTTP(res, req)

			s.Equal(tc.expectStatus, res.Code)
			body := res.Body.String()
			s.True(strings.HasPrefix(body, tc.expectBodyStartsWith))
		})
	}
}
