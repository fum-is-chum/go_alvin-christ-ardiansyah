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

// get all blogs
func GetBlogsController(c echo.Context) error {
	blogs, err := repositories.SelectBlogs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error on fetching blogs"))
	}

	var response = []models.BlogResponse{}
	for _, value := range blogs {
		response = append(response, models.BlogResponse{
			Id:        value.ID,
			Title:     value.Title,
			Content:   value.Content,
			UserID:    value.UserID,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get all blogs", response))
}

// get blog by id
func GetBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Bad Request: Id invalid"))
	}

	// your solution here
	blog, err := repositories.SelectBlogById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(fmt.Sprintf("Error on fetching blog with id %d", id)))
	}

	if blog.ID == 0 {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("Blog not found"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse(fmt.Sprintf("Sucess fetch blog with id %d", id), models.BlogResponse{
		Id:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		UserID:    blog.UserID,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}))
}

// create new blog

func CreateBlogController(c echo.Context) error {
	var newBlog models.Blog
	bindErr := c.Bind(&newBlog)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(fmt.Sprintf("Error bind %s", bindErr.Error())))
	}

	err := repositories.InsertBlog(newBlog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(fmt.Sprintf("Error insert blog: %s", err.Error())))
	}

	return c.JSON(http.StatusCreated, helpers.SuccessResponse("Success create blog"))
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Bad Request: Id invalid"))
	}

	err = repositories.DeleteBlog(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Delete blog failed"))
	}

	return c.String(http.StatusNoContent, "")
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Bad Request: Id invalid"))
	}

	var updateBlog models.Blog
	bindErr := c.Bind(&updateBlog)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(fmt.Sprintf("Error bind %s", bindErr.Error())))
	}

	updateBlog.ID = uint(id)
	err = repositories.UpdateBlog(updateBlog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Update blog failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success update blog"))
}
