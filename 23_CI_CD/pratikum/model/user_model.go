package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"not null;check:name <> ''"`
	Email    string `json:"email" form:"email" gorm:"unique;not null;check:email <> ''"`
	Password string `json:"password" form:"password" gorm:"not null;check:password <> ''"`
}
