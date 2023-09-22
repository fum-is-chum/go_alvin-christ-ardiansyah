package routes

import (
	"soal-prioritas-2/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// user route
	user := e.Group("/users")
	// Route / to handler function
	user.GET("", controllers.GetUsersController)
	user.GET("/:id", controllers.GetUserController)
	user.POST("", controllers.CreateUserController)
	user.DELETE("/:id", controllers.DeleteUserController)
	user.PUT("/:id", controllers.UpdateUserController)

	// book route
	book := e.Group("/books")
	book.GET("", controllers.GetBooksController)
	book.GET("/:id", controllers.GetBookController)
	book.POST("", controllers.CreateBookController)
	book.DELETE("/:id", controllers.DeleteBookController)
	book.PUT("/:id", controllers.UpdateBookController)
	return e
}
