package routes

import (
	"belajar-go-echo/controller"
	m "belajar-go-echo/middlewares"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	// middlewares
	m.LoggerMiddleware(e)

	// user
	userRepository := repository.NewUserRepository(db)
	userService := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userService)

	// auth
	authService := usecase.NewAuthUseCase(userRepository)
	authController := controller.NewAuthController(authService)

	userGroup := e.Group("/users")
	userGroup.GET("", userController.GetUser)
	userGroup.POST("", userController.Create)
	
	e.POST("/login", authController.Login)
}
