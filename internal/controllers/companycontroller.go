package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/fercarcedo/quadrant-server/internal/services"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type CompanyController struct {
	service services.CompanyService
}

func NewCompanyController(service services.CompanyService) *CompanyController {
	return &CompanyController { service }
}

func (controller *CompanyController) CreateCompany(c *gin.Context) {
	var company models.Company
	c.BindJSON(&company)
	if err := controller.service.CreateCompany(&company); err != nil {
		c.JSON(400, gin.H{})
	} else {
		c.JSON(201, gin.H{})
	}
}