package repositories

import (
	"errors"
	"restful-api-testing/config"
	"restful-api-testing/helpers"
	"restful-api-testing/middlewares"
	"restful-api-testing/models"
)

func CheckLogin(email string, password string) (models.User, string, error) {
	var data models.User
	
	// get user by email
	tx := config.DB.Where("email = ?", email).First(&data)
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

func SelectUsers() ([]models.User, error) {
	var users []models.User

	tx := config.DB.Order("created_at desc").Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func SelectUserById(Id uint) (models.User, error) {
	var user models.User
	tx := config.DB.First(&user, Id)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}

	return user, nil
}

func InsertUser(user models.User) error {
	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		return errors.New("Hash password failed")
	}
	user.Password = hash

	tx := config.DB.Create(&user)
	if tx.Error != nil {
		return errors.New("Insert user failed")
	}
	return nil
}

func DeleteUser(Id uint) error {
	tx := config.DB.Delete(&models.User{}, Id)
	if tx.Error != nil {
		return errors.New("Delete user failed")
	}
	return nil
}

func UpdateUser(user models.User) error {
	updatedUser := models.User{
		Name:  user.Name,
		Email: user.Email,
	}

	tx := config.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(updatedUser)
	if tx.Error != nil {
		return errors.New("Update user failed")
	}

	if tx.RowsAffected == 0 {
		return errors.New("User not found")
	}

	return nil
}
