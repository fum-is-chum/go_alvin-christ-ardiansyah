package controllers

import (
	"fmt"
	"latihan/helpers"
	"latihan/models"
	"latihan/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	var loginReq = models.LoginRequest{}
	errBind := c.Bind(&loginReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(fmt.Sprintf("Error bind: %s", errBind.Error())))
	}

	data, token, err := repositories.CheckLogin(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(fmt.Sprintf("Error login: %s", err.Error())))
	}

	response := map[string]any{
		"token":   token,
		"user_id": data.ID,
		"name":    data.Name,
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success login", response))
}
