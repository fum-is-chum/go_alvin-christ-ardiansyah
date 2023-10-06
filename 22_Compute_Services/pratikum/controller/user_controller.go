package controller

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/helpers"
	"belajar-go-echo/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userController struct {
	useCase usecase.UserUseCase
}

func NewUserController(userUsecase usecase.UserUseCase) *userController {
	return &userController{userUsecase}
}

func (u *userController) GetUser(c echo.Context) error {
	users, err := u.useCase.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(fmt.Sprintf("Error: %s", err.Error())))
	}

	var response []dto.UserResponse

	for _, value := range users {
		response = append(response, dto.UserResponse{
			Id: value.ID,
			Name: value.Name,
			Email: value.Email,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get users", response))
}

func (u *userController) Create(c echo.Context) error {
	var payload dto.CreateUserRequest

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(fmt.Sprintf("Error bind: %s", err.Error())))
	}

	// TODO: Validate whether all fields exist in request

	err := u.useCase.Create(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(fmt.Sprintf("Error: %s", err.Error())))
	}

	return c.JSON(http.StatusCreated, helpers.SuccessResponse("Success create user"))
}