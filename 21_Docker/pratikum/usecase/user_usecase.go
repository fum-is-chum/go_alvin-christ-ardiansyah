package usecase

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/helpers"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"errors"
)

type UserUseCase interface {
	GetUsers() ([]model.User, error)
	Create(payload dto.CreateUserRequest) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func createUserRequestToUser(data dto.CreateUserRequest) (model.User, error) {
	hash, err := helpers.HashPassword(data.Password)
	if err != nil {
		return model.User{}, errors.New("Hash password failed")
	}

    return model.User{
        Name:     data.Name,
        Email:    data.Email,
        Password: hash,
    }, nil
}

func NewUserUseCase(userRepo repository.UserRepository) *userUseCase {
	return &userUseCase{userRepo}
}

func (s *userUseCase) GetUsers() ([]model.User, error) {
	users, err := s.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userUseCase) Create(payload dto.CreateUserRequest) error {
	userData, err := createUserRequestToUser(payload)
	if err != nil {
		return err
	}
	if err := s.userRepository.Create(userData); err != nil {
		return err
	}

	return nil
}