package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/fercarcedo/quadrant-server/internal/services"
	"github.com/fercarcedo/quadrant-server/internal/models"
	"strconv"
)

type MachineInspectionController struct {
	service services.MachineInspectionService
}

func NewMachineInspectionController(service services.MachineInspectionService) *MachineInspectionController {
	return &MachineInspectionController { service }
}

func (controller *MachineInspectionController) GetMachineInspectionsByMachineId(c *gin.Context) {
	machineId, err := strconv.Atoi(c.Param("machineId"))
	onlyObservations := c.DefaultQuery("onlyObservations", "0") == "1"
	if err != nil {
		c.JSON(400, gin.H{})
	} else {
		machineInspections, err := controller.service.GetMachineInspectionsByMachineId(machineId, onlyObservations)
		if err != nil {
			c.JSON(400, gin.H{})
		} else {
			c.JSON(200, machineInspections)
		}
	}
}

func (controller *MachineInspectionController) CreateMachineInspection(c *gin.Context) {
	var machineInspection models.MachineInspection
	c.BindJSON(&machineInspection)
	machineId, err := strconv.Atoi(c.Param("machineId"))
	if err != nil {
		c.JSON(400, gin.H{})
	} else {
		machineInspection.MachineId = machineId
		if err := controller.service.CreateMachineInspection(&machineInspection); err != nil {
			c.JSON(400, gin.H{})
		} else {
			c.JSON(201, gin.H{})
		}
	}
}