package repository

import (
	"belajar-go-echo/helpers"
	"belajar-go-echo/model"
	"errors"

	mock "github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func NewMockUserRepo() *mockUserRepository {
	return &mockUserRepository{}
}

func (m *mockUserRepository) GetUsers() ([]model.User, error) {
	ret := m.Called()
	expectedUsers := []model.User{
		{Name: "User1", Email: "user1@example.com", Password: "password1"},
		{Name: "User2", Email: "user2@example.com", Password: "password2"},
	}

	return expectedUsers, ret.Error(1)
}

func (m *mockUserRepository) Create(data model.User) error {
	ret := m.Called(data)
	return ret.Error(0)
}

func (m *mockUserRepository) GetUserByEmail(email string) (model.User, error) {
	ret := m.Called(email)
	hash, err := helpers.HashPassword("123")
	if err != nil {
		return model.User{}, err
	}
	expectedUser := model.User{
		Name:     "Alvin",
		Email:    "alvinardiansyah2002@gmail.com",
		Password: hash,
	}

	if email != expectedUser.Email {
		return model.User{}, errors.New("User Not Found")
	}
	return expectedUser, ret.Error(1)
}
