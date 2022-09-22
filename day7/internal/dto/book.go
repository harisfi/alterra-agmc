package dto

import "github.com/harisfi/alterra-agmc/day7/internal/model"

type CreateBookRequest struct {
	Title     string     `json:"title" validate:"required"`
	Author    model.User `json:"author" validate:"required"`
	Publisher string     `json:"publisher"`
}

type UpdateBookRequest struct {
	Title     *string     `json:"title"`
	Author    *model.User `json:"author"`
	Publisher *string     `json:"publisher"`
}
