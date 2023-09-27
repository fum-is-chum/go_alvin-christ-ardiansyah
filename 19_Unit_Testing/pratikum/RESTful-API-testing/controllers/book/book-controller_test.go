package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"restful-api-testing/config"
	"restful-api-testing/models"
	"restful-api-testing/repositories"
	"restful-api-testing/utils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BookResponse struct {
	Message string
	Data    []models.Book
}

// ------- Utility ----------
func CheckTestBookExists() bool {
	tx := config.DB.First(&models.Book{}).Where("title = ?", "alta")
	if tx.Error != nil {
		return false
	}

	return true
}

func InsertDataBookForBookController(br *repositories.BookRepository) error {

	// check if book exist
	exist := CheckTestBookExists()

	if !exist {
		book := models.Book{
			Title:     "alta",
			Author:    "Alterra",
			Publisher: "Alterra",
		}

		err := br.InsertBook(book)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestGetBooksController(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		expectedCode int
		sizeData     int
	}{
		{
			name:         "Get Books",
			path:         "/books",
			expectedCode: http.StatusOK,
			sizeData:     1,
		},
	}

	e, db := utils.InitEchoTestAPI()
	bookRepo := repositories.NewBookRepository(db)
	bookController := NewBookController(*bookRepo)

	err := InsertDataBookForBookController(bookRepo)
	assert.Equal(t, nil, err)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Create a request with the specific book ID
			req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, bookController.GetBooksController(c)) {
				body := rec.Body.String()
				var response BookResponse

				assert.Equal(t, testCase.expectedCode, rec.Code, body)

				err := json.Unmarshal([]byte(body), &response)

				if err != nil {
					assert.Error(t, err, "Error")
				}

				assert.Equal(t, testCase.sizeData, len(response.Data))
			}
		})
	}
}

func TestGetBookController(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		id           string
		expectedCode int
	}{
		{
			name:         "Get Book with valid ID",
			path:         "/books",
			id:           "1",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid Get Book with non-number ID",
			path:         "/books",
			id:           "abcd",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Get Book with non-existing ID",
			path:         "/books",
			id:           "123",
			expectedCode: http.StatusNotFound,
		},
	}

	e, db := utils.InitEchoTestAPI()
	bookRepo := repositories.NewBookRepository(db)
	bookController := NewBookController(*bookRepo)

	err := InsertDataBookForBookController(bookRepo)
	assert.Equal(t, nil, err)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Create a request with the specific book ID
			req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(testCase.id)

			if assert.NoError(t, bookController.GetBookController(c)) {
				body := rec.Body.String()
				assert.Equal(t, testCase.expectedCode, rec.Code, body)
			}
		})
	}
}

func TestCreateBook(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		expectedCode int
		payload      models.Book
	}{
		{
			name:         "Create Book with valid payload",
			path:         "/books",
			expectedCode: http.StatusCreated,
			payload: models.Book{
				Title:     "Harry Potter",
				Author:    "JK Rowling",
				Publisher: "Warner Bros",
			},
		},
		{
			name:         "Create Book with invalid payload",
			path:         "/books",
			expectedCode: http.StatusInternalServerError,
			payload: models.Book{
				Title:  "Sherlock Holmes",
				Author: "Digital",
			},
		},
	}

	e, db := utils.InitEchoTestAPI()
	bookRepo := repositories.NewBookRepository(db)
	bookController := NewBookController(*bookRepo)

	err := InsertDataBookForBookController(bookRepo)
	assert.Equal(t, nil, err)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			payload, err := json.Marshal(testCase.payload)
			assert.Equal(t, err, nil)

			req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(string(payload)))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, bookController.CreateBookController(c)) {
				body := rec.Body.String()
				assert.Equal(t, testCase.expectedCode, rec.Code, body)

				// delete created record
				if rec.Code == http.StatusCreated {
					tx := config.DB.Unscoped().Where("title = ?", testCase.payload.Title).Delete(&models.Book{})
					assert.Equal(t, tx.Error, nil)
					assert.Equal(t, tx.RowsAffected, int64(1))
				}
			}
		})
	}
}

func TestUpdateBook(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		id           string
		expectedCode int
		payload      models.Book
	}{
		{
			name:         "Update Book with valid ID",
			path:         "/books",
			id:           "1",
			expectedCode: http.StatusOK,
			payload: models.Book{
				Title:     "alta_update",
				Author:    "Alterra",
				Publisher: "Alterra",
			},
		},
		{
			name:         "Update Book with non-number ID",
			path:         "/books",
			id:           "abc",
			expectedCode: http.StatusBadRequest,
			payload: models.Book{
				Title:     "alta_update",
				Author:    "Alterra",
				Publisher: "Alterra",
			},
		},
		{
			name:         "Update Book with not existing ID",
			path:         "/books",
			id:           "123",
			expectedCode: http.StatusBadRequest,
			payload: models.Book{
				Title:     "alta_update",
				Author:    "Alterra",
				Publisher: "Alterra",
			},
		},
	}

	e, db := utils.InitEchoTestAPI()
	bookRepo := repositories.NewBookRepository(db)
	bookController := NewBookController(*bookRepo)

	err := InsertDataBookForBookController(bookRepo)
	assert.Equal(t, nil, err)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			payload, err := json.Marshal(testCase.payload)
			assert.Equal(t, err, nil)

			req := httptest.NewRequest(http.MethodPut, testCase.path, strings.NewReader(string(payload)))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(testCase.id)

			if assert.NoError(t, bookController.UpdateBookController(c)) {
				body := rec.Body.String()
				assert.Equal(t, testCase.expectedCode, rec.Code, body)
			}
		})
	}
}

func TestDeleteBook(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		id           string
		expectedCode int
	}{
		{
			name:         "Delete Book with valid ID",
			path:         "/books",
			id:           "0",
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "Delete Book with non-number ID",
			path:         "/books",
			id:           "abc",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Delete Book with not existing ID",
			path:         "/books",
			id:           "123",
			expectedCode: http.StatusBadRequest,
		},
	}

	dummyBook := models.Book{
		Title:     "Harry Potter",
		Author:    "JK Rowling",
		Publisher: "Warner Bros",
	}

	e, db := utils.InitEchoTestAPI()
	bookRepo := repositories.NewBookRepository(db)
	bookController := NewBookController(*bookRepo)

	err := InsertDataBookForBookController(bookRepo)
	assert.Equal(t, nil, err)

	// insert dummy data
	payload, err := json.Marshal(dummyBook)
	assert.Equal(t, err, nil)

	dummy_req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(string(payload)))
	dummy_req.Header.Set("Content-Type", "application/json")
	dummy_rec := httptest.NewRecorder()
	c_dummy := e.NewContext(dummy_req, dummy_rec)

	var lastId int

	if assert.NoError(t, bookController.CreateBookController(c_dummy)) {
		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				// if expected is NoContent (valid), get last record ID's
				if testCase.expectedCode == http.StatusNoContent {
					tx := config.DB.Table("books").Select("id").Where("title = ?", dummyBook.Title).Find(&lastId)
					assert.Equal(t, tx.Error, nil)

					testCase.id = fmt.Sprintf("%d", lastId)
				}

				// test delete dummy data
				req := httptest.NewRequest(http.MethodDelete, testCase.path, nil)
				req.Header.Set("Content-Type", "application/json")
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)

				c.SetPath("/:id")
				c.SetParamNames("id")
				c.SetParamValues(testCase.id)

				if assert.NoError(t, bookController.DeleteBookController(c)) {
					assert.Equal(t, testCase.expectedCode, rec.Code)
				}
			})
		}
	}

	// delete book
	if dummy_rec.Code == http.StatusCreated {
		tx := config.DB.Unscoped().Where("id = ?", lastId).Delete(&models.Book{})
		assert.Equal(t, tx.Error, nil)
		assert.Equal(t, tx.RowsAffected, int64(1))
	}
}
