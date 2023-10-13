package controller

import (
	"fmt"
	"golang-ai/dto"
	responseHelper "golang-ai/helpers/response"
	"golang-ai/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type promptController struct {
	useCase usecase.PromptUseCase
}

func NewPromptController(promptUsecase usecase.PromptUseCase) *promptController {
	return &promptController{promptUsecase}
}

func (p *promptController) GetLaptopRecommendation(c echo.Context) error {
	var userRequest dto.LaptopSpecRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responseHelper.FailedResponse(fmt.Sprintf("Bad Request: %s", err.Error())))
	}

	answer, err := p.useCase.GetLaptopRecommendation(&userRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responseHelper.FailedResponse(fmt.Sprintf("Error: %s", err.Error())))
	}

	return c.JSON(http.StatusOK, responseHelper.SuccessWithDataResponse("Sucesss", answer))
}