package usecase

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUsers(t *testing.T) {
	mockRepo := repository.NewMockUserRepo()
	returnData := []model.User{
		{Name: "User1", Email: "user1@example.com", Password: "password1"},
		{Name: "User2", Email: "user2@example.com", Password: "password2"},
	}

	t.Run("Sucess Get Users", func(t *testing.T) {
		mockRepo.On("GetUsers", mock.Anything).Return(returnData, nil).Once()
		service := NewUserUseCase(mockRepo)

		res, err := service.GetUsers()
		assert.NoError(t, err)
		assert.Equal(t, len(returnData), len(res))
		assert.Equal(t, returnData[0].Name, res[0].Name)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Get Users", func(t *testing.T) {
        expectedErr := errors.New("database error") // Define the expected error
        mockRepo.On("GetUsers", mock.Anything).Return(nil, expectedErr).Once()
        service := NewUserUseCase(mockRepo)

        res, err := service.GetUsers()
        assert.Error(t, err)
        assert.Nil(t, res) // No users should be returned in case of an error
        assert.Equal(t, expectedErr, err)
        mockRepo.AssertExpectations(t)
    })
}

func TestCreateUser(t *testing.T) {
	t.Run("Success create user", func(t *testing.T) {
		data := dto.CreateUserRequest{
			Name:     "Alvin",
			Email:    "alvinardiansyah2002@gmail.com",
			Password: "123",
		}
	
		user,err := createUserRequestToUser(data)
		assert.Equal(t, err, nil)

		mockRepo := repository.NewMockUserRepo()
		mockRepo.On("Create", mock.MatchedBy(func(u model.User) bool {
			return u.Name == user.Name &&
				   u.Email == user.Email
		})).Return(nil)
	
		service := NewUserUseCase(mockRepo)
		err = service.Create(data)
	
		assert.Equal(t, err, nil)
	
		// Use userData when asserting the Create method is called
		mockRepo.AssertExpectations(t)
	})

	t.Run("Failed Create User", func(t *testing.T) {
		data := dto.CreateUserRequest{
			Name:     "Alvin",
			Email:    "alvinardiansyah2002@gmail.com",
			Password: "",
		}

		mockRepo := repository.NewMockUserRepo()
		mockRepo.On("Create", mock.Anything).Return(errors.New("Password is empty"))
		service := NewUserUseCase(mockRepo)
		err := service.Create(data)

		assert.NotEqual(t, err, nil)
		mockRepo.AssertExpectations(t)
	})
}
