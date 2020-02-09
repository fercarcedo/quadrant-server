package services_test

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"github.com/fercarcedo/quadrant-server/internal/models"
	"github.com/fercarcedo/quadrant-server/internal/services"
)

type mockUserDAO struct {
	mock.Mock
}

func (m *mockUserDAO) GetUsersByCompanyId(companyId int) ([]models.User, error) {
	args := m.Called(companyId)
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *mockUserDAO) CreateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserDAO) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserDAO) DeleteUser(userId int) error {
	args := m.Called(userId)
	return args.Error(0)
}

func TestGetUsersByCompanyId(t *testing.T) {
	dao := new(mockUserDAO)
	users := []models.User {
		models.User{ Id: 1, Name: "firstuser", IsAdmin: true, CompanyId: 2 },
		models.User{ Id: 2, Name: "seconduser", IsAdmin: false, CompanyId: 2 },
	}
	dao.On("GetUsersByCompanyId", 2).Return(users, nil)
	service := services.NewUserService(dao)
	usersByCompanyId, err := service.GetUsersByCompanyId(2)
	assert.Nil(t, err)
	assert.Equal(t, users, usersByCompanyId)
	dao.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	dao := new(mockUserDAO)
	user := &models.User{ 
		Id: 1, 
		Name: "firstuser", 
		IsAdmin: false, 
		CompanyId: 2,
	}
	dao.On("CreateUser", user).Return(nil)
	service := services.NewUserService(dao)
	assert.Nil(t, service.CreateUser(user))
	dao.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	dao := new(mockUserDAO)
	user := &models.User{ 
		Id: 1, 
		Name: "firstuser", 
		IsAdmin: false, 
		CompanyId: 2,
	}
	dao.On("UpdateUser", user).Return(nil)
	service := services.NewUserService(dao)
	assert.Nil(t, service.UpdateUser(user))
	dao.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	dao := new(mockUserDAO)
	dao.On("DeleteUser", 4).Return(nil)
	service := services.NewUserService(dao)
	assert.Nil(t, service.DeleteUser(4))
	dao.AssertExpectations(t)
}