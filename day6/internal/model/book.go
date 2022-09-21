package model

type Book struct {
	IDModel
	Title     string `json:"title" form:"title" validate:"required"`
	AuthorId  uint   `json:"author" form:"author" validate:"required"`
	Author    User   `json:"-" form:"-"`
	Publisher string `json:"publisher" form:"publisher" validate:"required"`
	TimestampModel
}
