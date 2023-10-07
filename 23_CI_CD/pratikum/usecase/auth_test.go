package usecase

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/helpers"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	mockRepo := repository.NewMockUserRepo()

	t.Run("Success Login", func(t *testing.T) {
		authRequest := dto.AuthRequest{
			Email:    "alvinardiansyah2002@gmail.com",
			Password: "123",
		}
		hash, err := helpers.HashPassword(authRequest.Password)
		assert.Equal(t, err, nil)
	
		expectedUser := model.User{
			Name: "Alvin",
			Email: authRequest.Email,
			Password: hash,
		}
		
		mockRepo.On("GetUserByEmail", authRequest.Email).Return(expectedUser, nil)

		service := NewAuthUseCase(mockRepo)

		// test login
		_, token, err := service.Login(authRequest)

		assert.Equal(t, err, nil)
		assert.NotEqual(t, token, nil)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Wrong password", func(t *testing.T) {
		authRequest := dto.AuthRequest{
			Email:    "alvinardiansyah2002@gmail.com",
			Password: "12345",
		}
		hash, err := helpers.HashPassword(authRequest.Password)
		assert.Equal(t, err, nil)
	
		expectedUser := model.User{
			Name: "Alvin",
			Email: authRequest.Email,
			Password: hash,
		}

		mockRepo.On("GetUserByEmail", authRequest.Email).Return(expectedUser, nil)
		service := NewAuthUseCase(mockRepo)

		// test login
		_, token, err := service.Login(authRequest)

		assert.NotEqual(t, err, nil)
		assert.Equal(t, token, "")
	})

	t.Run("User Not Found", func(t *testing.T){
		authRequest := dto.AuthRequest{
			Email:    "asdf@gmail.com",
			Password: "12345",
		}
		hash, err := helpers.HashPassword(authRequest.Password)
		assert.Equal(t, err, nil)
	
		expectedUser := model.User{
			Name: "Alvin",
			Email: authRequest.Email,
			Password: hash,
		}

		mockRepo.On("GetUserByEmail", authRequest.Email).Return(expectedUser, nil)
		service := NewAuthUseCase(mockRepo)

		// test login
		_, token, err := service.Login(authRequest)
		assert.Equal(t, err, nil)
		assert.Equal(t, token, "")
	})
}