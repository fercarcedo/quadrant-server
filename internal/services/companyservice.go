package services

import (
	"github.com/fercarcedo/quadrant-server/internal/daos"
	"github.com/fercarcedo/quadrant-server/internal/models"
	"github.com/google/uuid"
)

type CompanyService interface {
	CreateCompany(company *models.Company) error
}

type companyService struct {
	dao daos.CompanyDAO
}

func NewCompanyService(dao daos.CompanyDAO) CompanyService {
	return &companyService{ dao }
}

func (service *companyService) CreateCompany(company *models.Company) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	company.Code = id.String()
	return service.dao.CreateCompany(company)
}