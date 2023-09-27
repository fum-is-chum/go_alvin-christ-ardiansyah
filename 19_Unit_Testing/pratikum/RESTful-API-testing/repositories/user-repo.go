package repositories

import (
	"errors"
	"restful-api-testing/helpers"
	"restful-api-testing/middlewares"
	"restful-api-testing/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CheckLogin(email string, password string) (models.User, string, error) {
	var data models.User

	// get user by email
	tx := ur.db.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return models.User{}, "", tx.Error
	}

	// verify password
	_, err := helpers.VerifyPassword(data.Password, password)
	if err != nil {
		return models.User{}, "", err
	}

	var token string
	if tx.RowsAffected > 0 {
		var errToken error
		token, errToken = middlewares.CreateToken(int(data.ID))
		if errToken != nil {
			return models.User{}, "", errToken
		}
	}
	return data, token, nil
}

func (ur *UserRepository) SelectUsers() ([]models.User, error) {
	var users []models.User

	tx := ur.db.Order("created_at desc").Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (ur *UserRepository) SelectUserById(Id uint) (models.User, error) {
	var user models.User
	tx := ur.db.First(&user, Id)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}

	return user, nil
}

func (ur *UserRepository) InsertUser(user models.User) error {
	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		return errors.New("Hash password failed")
	}
	user.Password = hash

	tx := ur.db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ur *UserRepository) DeleteUser(Id uint) error {
	tx := ur.db.Delete(&models.User{}, Id)
	if tx.Error != nil {
		return errors.New("Delete user failed")
	}
	
	if tx.RowsAffected == 0 {
		return errors.New("User not found")
	}

	return nil
}

func (ur *UserRepository) UpdateUser(user models.User) error {
	updatedUser := models.User{
		Name:  user.Name,
		Email: user.Email,
	}

	tx := ur.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(updatedUser)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("User not found")
	}

	return nil
}
