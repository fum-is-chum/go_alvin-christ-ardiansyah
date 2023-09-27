package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"restful-api-testing/helpers"
	"restful-api-testing/models"
	"restful-api-testing/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	repo repositories.UserRepository
}

func NewUserController(ur repositories.UserRepository) *UserController {
	return &UserController{
		repo: ur,
	}
}

// get all users
func (uc *UserController) GetUsersController(c echo.Context) error {
	users, err := uc.repo.SelectUsers()
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
func (uc *UserController) GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Bad Request: Id invalid"))
	}

	// your solution here
	user, err := uc.repo.SelectUserById(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, helpers.FailedResponse("User not found"))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(fmt.Sprintf("Error on fetching user with id %d", id)))
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse(fmt.Sprintf("Sucess fetch user with id %d", id), models.UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}))
}

// create new user
func (uc *UserController) CreateUserController(c echo.Context) error {
	var newUser models.User
	bindErr := c.Bind(&newUser)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(fmt.Sprintf("Error bind %s", bindErr.Error())))
	}

	err := uc.repo.InsertUser(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error insert user"))
	}

	return c.JSON(http.StatusCreated, helpers.SuccessResponse("Success create user"))
}

// delete user by id
func (uc *UserController) DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Bad Request: Id invalid"))
	}

	err = uc.repo.DeleteUser(uint(id))
	if err != nil {
		if err.Error() == "User not found" {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Delete user failed"))
	}

	return c.String(http.StatusNoContent, "")
}

// update user by id
func (uc *UserController) UpdateUserController(c echo.Context) error {
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
	err = uc.repo.UpdateUser(updateUser)
	if err != nil {
		if err.Error() == "User not found" {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(fmt.Sprintf("Update user failed: %s", err.Error())))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success update user"))
}
