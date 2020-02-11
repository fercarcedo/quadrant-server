package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/fercarcedo/quadrant-server/internal/services"
	"github.com/fercarcedo/quadrant-server/internal/models"
	"strconv"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController { service }
}

func (controller *UserController) GetUsersByCompanyId(c *gin.Context) {
	companyId, err := strconv.Atoi(c.Param("companyId"))
	if err != nil {
		c.JSON(400, gin.H{})
	} else {
		users, err := controller.service.GetUsersByCompanyId(companyId)
		if err != nil {
			c.JSON(400, gin.H{})
		} else {
			c.JSON(200, users)
		}
	}
}

func (controller *UserController) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	if err := controller.service.CreateUser(&user); err != nil {
		c.JSON(400, gin.H{})
	} else {
		c.JSON(201, gin.H{})
	}
}

func (controller *UserController) UpdateUser(c *gin.Context) {
	inspectorId, err := strconv.Atoi(c.Param("userId"))
	var user models.User
	c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{})
	} else {
		user.Id = inspectorId
		if err := controller.service.UpdateUser(&user); err != nil {
			c.JSON(400, gin.H{})
		} else {
			c.JSON(204, gin.H{})
		}
	}
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	inspectorId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(400, gin.H{})
	} else {
		if err := controller.service.DeleteUser(inspectorId); err != nil {
			c.JSON(400, gin.H{})
		} else {
			c.JSON(204, gin.H{})
		}
	}
}



