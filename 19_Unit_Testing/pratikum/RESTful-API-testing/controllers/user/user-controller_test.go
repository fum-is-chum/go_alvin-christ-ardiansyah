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

type UserResponse struct {
	Message string
	Data    []models.User
}

// ------- Utility -----------
func CheckTestUserExist() bool {
	tx := config.DB.First(&models.User{}).Where("name = ?", "alta")
	if tx.Error != nil {
		return false
	}

	return true
}

func InsertDataUserForUserController(ur *repositories.UserRepository) error {

	// check if user exist
	exist:= CheckTestUserExist()

	if !exist {
		user := models.User{
			Name:     "alta",
			Password: "123",
			Email:    "alta@gmail.com",
		}

		err := ur.InsertUser(user)
		if err != nil {
			return err
		}
	}

	return nil
}

// ------ Unit Test --------
func TestGetUsersController(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		expectedCode int
		sizeData     int
	}{
		{
			name:         "Get Users",
			path:         "/users",
			expectedCode: http.StatusOK,
			sizeData:     1,
		},
	}

	e, db := utils.InitEchoTestAPI()
	userRepo := repositories.NewUserRepository(db)
	userController := NewUserController(*userRepo)

	err := InsertDataUserForUserController(userRepo)
	assert.Equal(t, nil, err)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Create a request with the specific user ID
			req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, userController.GetUsersController(c)) {
				body := rec.Body.String()
				var response UserResponse

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

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		id           string
		expectedCode int
	}{
		{
			name:         "Get User with valid ID",
			path:         "/users",
			id:           "1",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid Get User with non-number ID",
			path:         "/users",
			id:           "abcd",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Get User with non-existing ID",
			path:         "/users",
			id:           "123",
			expectedCode: http.StatusNotFound,
		},
	}

	e, db := utils.InitEchoTestAPI()
	userRepo := repositories.NewUserRepository(db)
	userController := NewUserController(*userRepo)

	err := InsertDataUserForUserController(userRepo)
	assert.Equal(t, nil, err)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Create a request with the specific user ID
			req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(testCase.id)

			if assert.NoError(t, userController.GetUserController(c)) {
				body := rec.Body.String()
				assert.Equal(t, testCase.expectedCode, rec.Code, body)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		expectedCode int
		payload      models.User
	}{
		{
			name:         "Create User with valid payload",
			path:         "/users",
			expectedCode: http.StatusCreated,
			payload: models.User{
				Name:     "Alvin",
				Email:    "alvinardiansyah2002@gmail.com",
				Password: "123",
			},
		},
		{
			name:         "Create User with invalid payload",
			path:         "/users",
			expectedCode: http.StatusInternalServerError,
			payload: models.User{
				Name:     "Alvin",
				Password: "123",
			},
		},
	}

	e, db := utils.InitEchoTestAPI()
	userRepo := repositories.NewUserRepository(db)
	userController := NewUserController(*userRepo)

	err := InsertDataUserForUserController(userRepo)
	assert.Equal(t, nil, err)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			payload, err := json.Marshal(testCase.payload)
			assert.Equal(t, err, nil)

			req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(string(payload)))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, userController.CreateUserController(c)) {
				body := rec.Body.String()
				assert.Equal(t, testCase.expectedCode, rec.Code, body)

				// delete created record
				if rec.Code == http.StatusCreated {
					tx := config.DB.Unscoped().Where("email = ?", testCase.payload.Email).Delete(&models.User{})
					assert.Equal(t, tx.Error, nil)
					assert.Equal(t, tx.RowsAffected, int64(1))
				}
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		id           string
		expectedCode int
		payload      models.User
	}{
		{
			name:         "Update User with valid ID",
			path:         "/users",
			id:           "1",
			expectedCode: http.StatusOK,
			payload: models.User{
				Name:     "alta_update",
				Email:    "alta@gmail.com",
				Password: "123",
			},
		},
		{
			name:         "Update User with non-number ID",
			path:         "/users",
			id:           "abc",
			expectedCode: http.StatusBadRequest,
			payload: models.User{
				Name:     "alta_update",
				Email:    "alta@gmail.com",
				Password: "123",
			},
		},
		{
			name:         "Update User with not existing ID",
			path:         "/users",
			id:           "123",
			expectedCode: http.StatusBadRequest,
			payload: models.User{
				Name:     "alta_update",
				Email:    "alta@gmail.com",
				Password: "123",
			},
		},
	}

	e, db := utils.InitEchoTestAPI()
	userRepo := repositories.NewUserRepository(db)
	userController := NewUserController(*userRepo)

	err := InsertDataUserForUserController(userRepo)
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

			if assert.NoError(t, userController.UpdateUserController(c)) {
				body := rec.Body.String()
				assert.Equal(t, testCase.expectedCode, rec.Code, body)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	var testCases = []struct {
		name         string
		path         string
		id           string
		expectedCode int
	}{
		{
			name:         "Delete User with valid ID",
			path:         "/users",
			id:           "0",
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "Delete User with non-number ID",
			path:         "/users",
			id:           "abc",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Delete User with not existing ID",
			path:         "/users",
			id:           "123",
			expectedCode: http.StatusBadRequest,
		},
	}

	dummyUser := models.User{
		Name:     "Alvin",
		Email:    "alvinardiansyah2002@gmail.com",
		Password: "123",
	}

	e, db := utils.InitEchoTestAPI()
	userRepo := repositories.NewUserRepository(db)
	userController := NewUserController(*userRepo)

	err := InsertDataUserForUserController(userRepo)
	assert.Equal(t, nil, err)

	// insert dummy data
	payload, err := json.Marshal(dummyUser)
	assert.Equal(t, err, nil)

	dummy_req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(payload)))
	dummy_req.Header.Set("Content-Type", "application/json")
	dummy_rec := httptest.NewRecorder()
	c_dummy := e.NewContext(dummy_req, dummy_rec)

	var lastId int

	if assert.NoError(t, userController.CreateUserController(c_dummy)) {
		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				// if expected is NoContent (valid), get last record ID's
				if testCase.expectedCode == http.StatusNoContent {
					tx := config.DB.Table("users").Select("id").Where("email = ?", dummyUser.Email).Find(&lastId)
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

				if assert.NoError(t, userController.DeleteUserController(c)) {
					assert.Equal(t, testCase.expectedCode, rec.Code)
				}
			})
		} 
	}

	// delete user
	if dummy_rec.Code == http.StatusCreated {
		tx := config.DB.Unscoped().Where("id = ?", lastId).Delete(&models.User{})
		assert.Equal(t, tx.Error, nil)
		assert.Equal(t, tx.RowsAffected, int64(1))
	}
}
