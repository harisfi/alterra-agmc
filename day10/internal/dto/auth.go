package dto

import "github.com/harisfi/alterra-agmc/day10/internal/model"

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}
type AuthLoginResponse struct {
	Token string `json:"token"`
	model.User
}

type AuthRegisterRequest struct {
	model.User
}
type AuthRegisterResponse struct {
	model.User
}
