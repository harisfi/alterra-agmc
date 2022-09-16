package models

type User struct {
	IDModel
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,gte=8"`
	TimestampModel
}
