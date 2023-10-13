package routes

import (
	"golang-ai/config"
	"golang-ai/controller"
	"golang-ai/usecase"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	bot := config.NewBot()

	botUsecase := usecase.NewPromptUseCase(bot)
	botController := controller.NewPromptController(botUsecase)
	e.POST("/bot-api", botController.GetLaptopRecommendation)
}
