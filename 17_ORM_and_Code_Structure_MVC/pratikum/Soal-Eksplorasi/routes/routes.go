package routes

import (
	"soal-eksplorasi/controllers"

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

	// blog route
	blog := e.Group("/blogs")
	// Route / to handler function
	blog.GET("", controllers.GetBlogsController)
	blog.GET("/:id", controllers.GetBlogController)
	blog.POST("", controllers.CreateBlogController)
	blog.DELETE("/:id", controllers.DeleteBlogController)
	blog.PUT("/:id", controllers.UpdateBlogController)

	return e
}
