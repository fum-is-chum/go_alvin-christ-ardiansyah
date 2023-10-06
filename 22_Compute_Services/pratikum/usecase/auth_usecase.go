package usecase

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/helpers"
	"belajar-go-echo/middlewares"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type AuthUseCase interface {
	Login(payload dto.AuthRequest) (model.User, string, error)
}

type authUseCase struct {
	userRepo repository.UserRepository
}

func NewAuthUseCase(userRepo repository.UserRepository) *authUseCase {
	return &authUseCase{userRepo}
}

func (s *authUseCase) Login(payload dto.AuthRequest) (model.User, string, error) {
	userData, err := s.userRepo.GetUserByEmail(payload.Email)
	if err != nil {
		return model.User{}, "", nil
	}

	// verify password
	err = helpers.VerifyPassword(userData.Password, payload.Password)
	if err != nil {
		return model.User{}, "", err
	}

	// get jwt token
	token, err := middlewares.CreateToken(int(userData.ID))
	if err != nil {
		return model.User{}, "", err
	}

	return userData, token, nil
}