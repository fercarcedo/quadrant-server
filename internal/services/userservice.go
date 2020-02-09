package services

import (
	"github.com/fercarcedo/quadrant-server/internal/daos"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type UserService interface {
	GetUsersByCompanyId(companyId int) ([]models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(userId int) error
}

type userService struct {
	dao daos.UserDAO
}

func NewUserService(dao daos.UserDAO) UserService {
	return &userService{ dao }
}

func (service *userService) GetUsersByCompanyId(companyId int) ([]models.User, error) {
	return service.dao.GetUsersByCompanyId(companyId)
}

func (service *userService) CreateUser(user *models.User) error {
	return service.dao.CreateUser(user)
}

func (service *userService) UpdateUser(user *models.User) error {
	return service.dao.UpdateUser(user)
}

func (service *userService) DeleteUser(userId int) error {
	return service.dao.DeleteUser(userId)
}
