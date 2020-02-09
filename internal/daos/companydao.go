package daos

import (
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type CompanyDAO interface {
	CreateCompany(company *models.Company) error
}

type companyDAO struct {}

func NewCompanyDAO() CompanyDAO {
	return &companyDAO{}
}

func (dao *companyDAO) CreateCompany(company *models.Company) error {
	return config.Config.DB.Insert(company)
}