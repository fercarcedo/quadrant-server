package main

import (
	"fmt"
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/db"
	"github.com/fercarcedo/quadrant-server/internal/daos"
	"github.com/fercarcedo/quadrant-server/internal/services"
	"github.com/fercarcedo/quadrant-server/internal/controllers"
	"github.com/gin-contrib/static"
)

func main() {
	var err error
	if err = config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	config.Config.DB, err = db.ConnectToDb()
	if err != nil {
		panic(fmt.Errorf("error connecting to the db: %s", err))
	}
	defer config.Config.DB.Close()
	companyController := controllers.NewCompanyController(services.NewCompanyService(daos.NewCompanyDAO()))
	machineController := controllers.NewMachineController(services.NewMachineService(daos.NewMachineDAO()))
	machineInspectionController := controllers.NewMachineInspectionController(services.NewMachineInspectionService(daos.NewMachineInspectionDAO()))
	userController := controllers.NewUserController(services.NewUserService(daos.NewUserDAO()))
	r := controllers.NewRouter(companyController, machineController, machineInspectionController, userController).SetUpRouter()
	r.Use(static.Serve("/", static.LocalFile("/public", false)))
	r.Run(fmt.Sprintf(":%d", config.Config.ServerPort))
}