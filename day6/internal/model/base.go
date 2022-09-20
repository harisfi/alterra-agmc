package model

import "time"

type IDModel struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

type TimestampModel struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
