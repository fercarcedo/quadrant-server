package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"errors"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/fercarcedo/quadrant-server/internal/controllers"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type mockUserService struct {
	mock.Mock
}

func (m *mockUserService) GetUsersByCompanyId(companyId int) ([]models.User, error) {
	args := m.Called(companyId)
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *mockUserService) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserService) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserService) DeleteUser(userId int) error {
	args := m.Called(userId)
	return args.Error(0)
}

func TestGetUsersByCompanyId(t *testing.T) {
	userService := new(mockUserService)
	users := []models.User { 
		models.User { Name: "user1", IsAdmin: false, CompanyId: 2 }, 
		models.User { Name: "user2", IsAdmin: true, CompanyId: 2 }, 
	}
	userService.On("GetUsersByCompanyId", 2).Return(users, nil)
	userController := controllers.NewUserController(userService)
	router := controllers.NewRouter(nil, nil, nil, userController).SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/companies/2/users", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	jsonBytes, _ := json.Marshal(&users)
	assert.Equal(t, string(jsonBytes)+"\n", w.Body.String())
}

func TestCreateUser(t *testing.T) {
	userService := new(mockUserService)
	userService.On("CreateUser", mock.MatchedBy(func (user *models.User) bool {
		return user.Name == "testuser" &&
			user.IsAdmin == false &&
			user.CompanyId == 2
	})).Return(nil)
	user := &models.User {
		Name: "testuser",
		IsAdmin: false,
		CompanyId: 2,
	}
	userController := controllers.NewUserController(userService)
	router := controllers.NewRouter(nil, nil, nil, userController).SetUpRouter()
	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}

func TestCreateUserBadRequest(t *testing.T) {
	userService := new(mockUserService)
	userService.On("CreateUser", mock.MatchedBy(func (user *models.User) bool {
		return user.Name == "testuser" &&
			user.IsAdmin == false &&
			user.CompanyId == 2
	})).Return(errors.New("Error"))
	user := &models.User {
		Name: "testuser",
		IsAdmin: false,
		CompanyId: 2,
	}
	userController := controllers.NewUserController(userService)
	router := controllers.NewRouter(nil, nil, nil, userController).SetUpRouter()
	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}

func TestUpdateUser(t *testing.T) {
	userService := new(mockUserService)
	userService.On("UpdateUser", mock.MatchedBy(func (user *models.User) bool {
		return user.Id == 1 &&
			user.Name == "testuser" &&
			user.IsAdmin == false &&
			user.CompanyId == 2
	})).Return(nil)
	user := &models.User {
		Name: "testuser",
		IsAdmin: false,
		CompanyId: 2,
	}
	userController := controllers.NewUserController(userService)
	router := controllers.NewRouter(nil, nil, nil, userController).SetUpRouter()
	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(user)
	req, _ := http.NewRequest("PUT", "/api/v1/users/1", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestUpdateUserBadRequest(t *testing.T) {
	userService := new(mockUserService)
	userService.On("UpdateUser", mock.MatchedBy(func (user *models.User) bool {
		return user.Id == 1 &&
			user.Name == "testuser" &&
			user.IsAdmin == false &&
			user.CompanyId == 2
	})).Return(errors.New("Error"))
	user := &models.User {
		Name: "testuser",
		IsAdmin: false,
		CompanyId: 2,
	}
	userController := controllers.NewUserController(userService)
	router := controllers.NewRouter(nil, nil, nil, userController).SetUpRouter()
	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(user)
	req, _ := http.NewRequest("PUT", "/api/v1/users/1", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}

func TestDeleteUser(t *testing.T) {
	userService := new(mockUserService)
	userService.On("DeleteUser", 1).Return(nil)
	userController := controllers.NewUserController(userService)
	router := controllers.NewRouter(nil, nil, nil, userController).SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 204, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestDeleteUserBadRequest(t *testing.T) {
	userService := new(mockUserService)
	userService.On("DeleteUser", 1).Return(errors.New("Error"))
	userController := controllers.NewUserController(userService)
	router := controllers.NewRouter(nil, nil, nil, userController).SetUpRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}