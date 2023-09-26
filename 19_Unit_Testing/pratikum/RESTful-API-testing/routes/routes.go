package routes

import (
	"restful-api-testing/controllers"
	m "restful-api-testing/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	// middlewares
	m.LoggerMiddleware(e)
	
	// users route
	user := e.Group("/users")
	user.GET("", controllers.GetUsersController, m.JWTMiddleware())
	user.GET("/:id", controllers.GetUserController, m.JWTMiddleware())
	user.POST("", controllers.CreateUserController)
	user.DELETE("/:id", controllers.DeleteUserController, m.JWTMiddleware())
	user.PUT("/:id", controllers.UpdateUserController, m.JWTMiddleware())

	user.POST("/login", controllers.LoginController)

	// books route
	book := e.Group("/books")
	book.Use(m.JWTMiddleware())
	book.GET("", controllers.GetBooksController)
	book.GET("/:id", controllers.GetBookController)
	book.POST("", controllers.CreateBookController)
	book.DELETE("/:id", controllers.DeleteBookController)
	book.PUT("/:id", controllers.UpdateBookController)
}
