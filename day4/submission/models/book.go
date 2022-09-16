package models

type Book struct {
	IDModel
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	TimestampModel
}
