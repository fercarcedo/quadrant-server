package daos

import (
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type CompanyDAO struct {}

func NewCompanyDAO() *CompanyDAO {
	return &CompanyDAO{}
}

func (dao *CompanyDAO) CreateCompany(company *models.Company) error {
	return config.Config.DB.Insert(company)
}