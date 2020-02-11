package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/fercarcedo/quadrant-server/internal/services"
	"github.com/fercarcedo/quadrant-server/internal/models"
	"strconv"
)

type MachineController struct {
	service services.MachineService
}

func NewMachineController(service services.MachineService) *MachineController {
	return &MachineController { service }
}

func (controller *MachineController) GetMachinesByInspectorIdAndRevision(c *gin.Context) {
	inspectorId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(400, gin.H{})
	} else {
		machines, err := controller.service.GetMachinesByInspectorIdAndRevision(inspectorId)
		if err != nil {
			c.JSON(500, gin.H{})
		} else {
			c.JSON(200, machines)
		}
	}
}

func (controller *MachineController) CreateMachine(c *gin.Context) {
	var machine models.Machine
	c.BindJSON(&machine)
	if err := controller.service.CreateMachine(&machine); err != nil {
		c.JSON(400, gin.H{})
	} else {
		c.JSON(201, gin.H{})
	}
}

func (controller *MachineController) UpdateMachine(c *gin.Context) {
	var machine models.Machine
	c.BindJSON(&machine)
	machineId, err := strconv.Atoi(c.Param("machineId"))
	if err != nil {
		c.JSON(400, gin.H{})
	} else {
		machine.Id = machineId
		if err := controller.service.UpdateMachine(&machine); err != nil {
			c.JSON(400, gin.H{})
		} else {
			c.JSON(204, gin.H{})
		}
	}
}

func (controller *MachineController) DeleteMachine(c *gin.Context) {
	machineId, err := strconv.Atoi(c.Param("machineId"))
	if err != nil {
		c.JSON(400, gin.H{})
	} else {
		if err := controller.service.DeleteMachine(machineId); err != nil {
			c.JSON(400, gin.H{})
		} else {
			c.JSON(204, gin.H{})
		}
	}
}