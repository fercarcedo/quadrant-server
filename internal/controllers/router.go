package controllers

import (	
	"github.com/gin-gonic/gin"
)

type Router struct {
	companyController *CompanyController
	machineController *MachineController
	machineInspectionController *MachineInspectionController
	userController *UserController
}

func NewRouter(
	companyController *CompanyController, 
	machineController *MachineController, 
	machineInspectionController *MachineInspectionController, 
	userController *UserController,
) *Router {
	return &Router {
		companyController,
		machineController,
		machineInspectionController,
		userController,
	}
}

func (router *Router) SetUpRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1") 
	{
		v1.POST("/companies", router.companyController.CreateCompany)
		v1.GET("/users/:userId/machines", router.machineController.GetMachinesByInspectorIdAndRevision)
		v1.POST("/machines", router.machineController.CreateMachine)
		v1.PUT("/machines/:machineId", router.machineController.UpdateMachine)
		v1.DELETE("/machines/:machineId", router.machineController.DeleteMachine)
		v1.GET("/machines/:machineId/inspections", router.machineInspectionController.GetMachineInspectionsByMachineId)
		v1.POST("/machines/:machineId/inspections", router.machineInspectionController.CreateMachineInspection)
		v1.GET("/companies/:companyId/users", router.userController.GetUsersByCompanyId)
		v1.POST("/users", router.userController.CreateUser)
		v1.PUT("/users/:userId", router.userController.UpdateUser)
		v1.DELETE("/users/:userId", router.userController.DeleteUser)
	}
	return r
}