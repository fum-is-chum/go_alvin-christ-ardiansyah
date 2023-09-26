package controllers

import (
	"fmt"
	"net/http"
	"restful-api-testing/helpers"
	"restful-api-testing/models"
	"restful-api-testing/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	books, err := repositories.SelectBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error on fetching books"))
	}

	var response = []models.BookResponse{}
	for _, value := range books {
		response = append(response, models.BookResponse{
			Id:        value.ID,
			Title:     value.Title,
			Author:    value.Author,
			Publisher: value.Publisher,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success get all books", response))
}

// get book by id
func GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Bad Request: Id invalid"))
	}

	// your solution here
	book, err := repositories.SelectBookById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(fmt.Sprintf("Error on fetching book with id %d", id)))
	}

	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, helpers.FailedResponse("Book not found"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse(fmt.Sprintf("Sucess fetch book with id %d", id), models.BookResponse{
		Id:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}))
}

// create new book
func CreateBookController(c echo.Context) error {
	var newBook models.Book
	bindErr := c.Bind(&newBook)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(fmt.Sprintf("Error bind %s", bindErr.Error())))
	}

	err := repositories.InsertBook(newBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Error insert book"))
	}

	return c.JSON(http.StatusCreated, helpers.SuccessResponse("Success create book"))
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Bad Request: Id invalid"))
	}

	err = repositories.DeleteBook(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Delete book failed"))
	}

	return c.String(http.StatusNoContent, "")
}

// update book by id
func UpdateBookController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("Bad Request: Id invalid"))
	}

	var updateBook models.Book
	bindErr := c.Bind(&updateBook)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(fmt.Sprintf("Error bind %s", bindErr.Error())))
	}

	updateBook.ID = uint(id)
	err = repositories.UpdateBook(updateBook)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse("Update book failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("Success update book"))
}
