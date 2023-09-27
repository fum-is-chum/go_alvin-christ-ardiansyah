package routes

import (
	ac "restful-api-testing/controllers/auth"
	bc "restful-api-testing/controllers/book"
	uc "restful-api-testing/controllers/user"
	m "restful-api-testing/middlewares"
	"restful-api-testing/repositories"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	// middlewares
	m.LoggerMiddleware(e)

	// repositories
	userRepo := repositories.NewUserRepository(db)
	bookRepo := repositories.NewBookRepository(db)

	userController := uc.NewUserController(*userRepo)
	// users route
	user := e.Group("/users")
	user.GET("", userController.GetUsersController, m.JWTMiddleware())
	user.GET("/:id", userController.GetUserController, m.JWTMiddleware())
	user.POST("", userController.CreateUserController)
	user.DELETE("/:id", userController.DeleteUserController, m.JWTMiddleware())
	user.PUT("/:id", userController.UpdateUserController, m.JWTMiddleware())

	authController := ac.NewAuthController(*userRepo)
	user.POST("/login", authController.LoginController)

	bookController := bc.NewBookController(*bookRepo)
	// books route
	book := e.Group("/books")
	book.Use(m.JWTMiddleware())
	book.GET("", bookController.GetBooksController)
	book.GET("/:id", bookController.GetBookController)
	book.POST("", bookController.CreateBookController)
	book.DELETE("/:id", bookController.DeleteBookController)
	book.PUT("/:id", bookController.UpdateBookController)
}
