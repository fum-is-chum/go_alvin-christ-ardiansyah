package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title" gorm:"unique;not null;check:title <> ''"`
	Author    string `json:"author" form:"author" gorm:"not null;check:author <> ''"`
	Publisher string `json:"publisher" form:"publisher" gorm:"not null;check:publisher <>''"`
}

type BookResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Publisher string    `json:"publisher"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
