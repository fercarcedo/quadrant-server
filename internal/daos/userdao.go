package daos

import (
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type UserDAO struct {}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) GetUsersByCompanyId(companyId int) ([]models.User, error) {
	var users []models.User
	err := config.Config.DB.Model(&users).Where("company_id = ?", companyId).Order("name ASC").Select()
	return users, err
}

func (dao *UserDAO) CreateUser(user *models.User) error {
	return config.Config.DB.Insert(user)
}

func (dao *UserDAO) UpdateUser(user *models.User) error {
	return config.Config.DB.Update(user)
}

func (dao *UserDAO) DeleteUser(userId int) error {
	user := &models.User{ Id: userId }
	return config.Config.DB.Delete(user)
}
