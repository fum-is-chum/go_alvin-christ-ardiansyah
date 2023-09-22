package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// ------------------- Utility --------------------
func findUserById(id int) (*User, int) {
	for i, user := range users {
		if user.Id == id {
			return &user, i
		}
	}
	return nil, -1
}

// -------------------- controller --------------------
// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	user, _ := findUserById(id)

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "User not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	_, index := findUserById(id)

	if index == -1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "User not found",
		})
	}

	users = append(users[:index], users[index+1:]...)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	updatedUser := new(User)
	err = c.Bind(&updatedUser)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	user, index := findUserById(id)
	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "User not found",
		})
	}

	users[index] = *updatedUser

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sucess",
		"user":    updatedUser,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
