package controllers

import (
	"fmt"
	"net/http"
	"soal-eksplorasi/helpers"
	"soal-eksplorasi/models"
	"soal-eksplorasi/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	users, err := repositories.SelectUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error on fetching users"))
	}

	var response = []models.UserResponse{}
	for _, value := range users {
		response = append(response, models.UserResponse{
			Id:        value.ID,
			Name:      value.Name,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get all users", response))
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Bad Request: Id invalid"))
	}

	// your solution here
	user, err := repositories.SelectUserById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(fmt.Sprintf("Error on fetching user with id %d", id)))
	}

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("User not found"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse(fmt.Sprintf("Sucess fetch user with id %d", id), models.MapUserToUserResponse(user)))
}

// create new user

func CreateUserController(c echo.Context) error {
	var newUser models.User
	bindErr := c.Bind(&newUser)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(fmt.Sprintf("Error bind %s", bindErr.Error())))
	}

	err := repositories.InsertUser(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error insert user"))
	}

	return c.JSON(http.StatusCreated, helpers.SuccessResponse("Success create user"))
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Bad Request: Id invalid"))
	}

	err = repositories.DeleteUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(fmt.Sprintf("Delete user failed: %s", err.Error())))
	}

	return c.String(http.StatusNoContent, "")
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Bad Request: Id invalid"))
	}

	var updateUser models.User
	bindErr := c.Bind(&updateUser)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(fmt.Sprintf("Error bind %s", bindErr.Error())))
	}

	updateUser.ID = uint(id)
	err = repositories.UpdateUser(updateUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Update user failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success update user"))
}
