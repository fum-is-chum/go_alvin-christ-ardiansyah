package repository

import (
	"belajar-go-echo/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	Create(data model.User) error
	GetUserByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db} 
}

func (u *userRepository) GetUsers() ([]model.User, error) {
	var users []model.User

	tx := u.db.Order("created_at desc").Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (u *userRepository) Create(data model.User) error {
	var userData model.User

	tx := u.db.Where("email = ?", data.Email).First(&userData)
	if tx.Error == nil {
		return errors.New("Duplicate email")
	}

	if err := u.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUserByEmail(email string) (model.User, error) {
	var userData model.User

	tx := u.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil{
		return model.User{}, tx.Error
	}

	return userData, nil
}