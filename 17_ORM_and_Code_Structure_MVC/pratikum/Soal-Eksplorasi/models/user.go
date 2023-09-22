package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Blogs    []Blog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserResponse struct {
	Id        uint           `json:"id"`
	Name      string         `json:"name,omitempty"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Blogs     []BlogResponse `json:"blogs"`
}

func MapUserToUserResponse(user User) UserResponse {
	userResponse := UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	// Map blogs from []Blog to []BlogResponse
	for _, blog := range user.Blogs {
		userResponse.Blogs = append(userResponse.Blogs, BlogResponse{
			Id:        blog.ID,
			Title:     blog.Title,
			Content:   blog.Content,
			UserID:    blog.UserID,
			CreatedAt: blog.CreatedAt,
			UpdatedAt: blog.UpdatedAt,
		})
	}

	return userResponse
}
