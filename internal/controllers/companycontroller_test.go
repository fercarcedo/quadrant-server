package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/fercarcedo/quadrant-server/internal/controllers"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type mockCompanyService struct {
	mock.Mock
}

func (m *mockCompanyService) CreateCompany(company *models.Company) error {
	args := m.Called(company)
	return args.Error(0)
}

func TestCreateCompany(t *testing.T) {
	companyService := new(mockCompanyService)
	companyService.On("CreateCompany", mock.MatchedBy(func(company *models.Company) bool {
		return company.Name == "testcompany"
	})).Return(nil)
	companyController := controllers.NewCompanyController(companyService)
	router := controllers.NewRouter(companyController, nil, nil, nil).SetUpRouter()
	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(&models.Company { Name: "testcompany" })
	req, _ := http.NewRequest("POST", "/api/v1/companies", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}

func TestCreateCompanyBadRequest(t *testing.T) {
	companyService := new(mockCompanyService)
	companyService.On("CreateCompany", mock.MatchedBy(func(company *models.Company) bool {
		return company.Name == "testcompany"
	})).Return(errors.New("Error"))	
	companyController := controllers.NewCompanyController(companyService)
	router := controllers.NewRouter(companyController, nil, nil, nil).SetUpRouter()
	w := httptest.NewRecorder()
	jsonStr, _ := json.Marshal(&models.Company { Name: "testcompany" })
	req, _ := http.NewRequest("POST", "/api/v1/companies", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{}\n", w.Body.String())
}