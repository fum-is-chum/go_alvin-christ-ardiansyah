package routes

import (
	"latihan/controllers"
	m "latihan/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

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
	return e
}
