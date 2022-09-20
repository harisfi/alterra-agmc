package model

type Book struct {
	IDModel
	Title     string `json:"title" form:"title" validate:"required"`
	Author    User   `json:"author" form:"author" validate:"required"`
	Publisher string `json:"publisher" form:"publisher" validate:"required"`
	TimestampModel
}
