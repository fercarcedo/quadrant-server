package daos

import (
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type UserDAO interface {
	GetUsersByCompanyId(companyId int) ([]models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(userId int) error
}

type userDAO struct {}

func NewUserDAO() UserDAO {
	return &userDAO{}
}

func (dao *userDAO) GetUsersByCompanyId(companyId int) ([]models.User, error) {
	var users []models.User
	err := config.Config.DB.Model(&users).Where("company_id = ?", companyId).Order("name ASC").Select()
	return users, err
}

func (dao *userDAO) CreateUser(user *models.User) error {
	return config.Config.DB.Insert(user)
}

func (dao *userDAO) UpdateUser(user *models.User) error {
	return config.Config.DB.Update(user)
}

func (dao *userDAO) DeleteUser(userId int) error {
	user := &models.User{ Id: userId }
	return config.Config.DB.Delete(user)
}
