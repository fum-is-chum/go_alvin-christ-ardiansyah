package controller

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/helpers"
	"belajar-go-echo/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authController struct {
	useCase usecase.AuthUseCase
}

func NewAuthController(authUsecase usecase.AuthUseCase) *authController {
	return &authController{authUsecase}
}

func (u *authController) Login(c echo.Context) error {
	var payload dto.AuthRequest

	if err := c.Bind(&payload); err != nil {
		return err
	}

	userData, token, err := u.useCase.Login(payload)
	if err != nil {
		return err
	}

	userResponse := dto.UserResponse{
		Id:        userData.ID,
		Name:      userData.Name,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	response := map[string]interface{}{
		"token": token,
		"user":  userResponse,
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Login success", response))
}
