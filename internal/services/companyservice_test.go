package services_test

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"github.com/fercarcedo/quadrant-server/internal/models"
	"github.com/fercarcedo/quadrant-server/internal/services"
	"github.com/google/uuid"
)

type mockCompanyDAO struct {
	mock.Mock
}

func (m *mockCompanyDAO) CreateCompany(company *models.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

func TestCreateCompany(t *testing.T) {
	dao := new(mockCompanyDAO)
	dao.On("CreateCompany", mock.MatchedBy(func(company *models.Company) bool {
		_, err := uuid.Parse(company.Code)
		return company.Name == "testcompany" && err == nil
	})).Return(nil)
	service := services.NewCompanyService(dao)
	company := &models.Company{
		Name: "testcompany",
	}
	assert.Nil(t, service.CreateCompany(company))
	dao.AssertExpectations(t)
}