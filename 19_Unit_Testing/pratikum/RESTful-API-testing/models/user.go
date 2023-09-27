package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"not null;check:name <> ''"`
	Email    string `json:"email" form:"email" gorm:"unique;not null;check:email <> ''"`
	Password string `json:"password" form:"password" gorm:"not null;check:password <> ''"`
}

type UserResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
