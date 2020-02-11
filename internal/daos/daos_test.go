package daos_test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
	"os"
	"time"
	"testing"
	"github.com/ory/dockertest"
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/db"
	"github.com/fercarcedo/quadrant-server/internal/models"
	"github.com/fercarcedo/quadrant-server/internal/daos"
	"github.com/stretchr/testify/assert"
)

var dbURL string

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}
	resource, err := pool.Run("postgres", "9.6-alpine", []string{"POSTGRES_PASSWORD=secret", "POSTGRES_DB=db"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	if err := pool.Retry(func() error {
		var err error
		var database *sql.DB
		dbURL = fmt.Sprintf("postgres://postgres:secret@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), "db")
		database, err = sql.Open("postgres", dbURL)
		if err != nil {
			return err
		}
		err = database.Ping()
		if err == nil {
			defer database.Close()
		}
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	config.Config.DBURL = dbURL
	config.Config.DB, _ = db.ConnectToDb()
	defer config.Config.DB.Close()

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestCreateCompany(t *testing.T) {
	company := &models.Company { Name: "company", Code: "4bb1500c-2b05-475c-bff2-ffe93942c697" }
	dao := daos.NewCompanyDAO()
	err := dao.CreateCompany(company)
	assert.Nil(t, err)
	defer config.Config.DB.Delete(company)
	var companies []models.Company
	config.Config.DB.Model(&companies).Select()
	assert.Equal(t, 1, len(companies))
	assert.Equal(t, company.Name, companies[0].Name)
	assert.Equal(t, company.Code, companies[0].Code)
	assert.Equal(t, 1, companies[0].Id)
}

func TestGetMachinesByInspectorIdAndRevision(t *testing.T) {
	company := &models.Company { Name: "Company", Code: "4bb1500c-2b05-475c-bff2-ffe93942c697" }
	daos.NewCompanyDAO().CreateCompany(company)
	defer config.Config.DB.Delete(company)
	var companies []models.Company
	config.Config.DB.Model(&companies).Select()
	inspector := &models.User { Name: "testuser", IsAdmin: false, CompanyId: companies[0].Id }
	daos.NewUserDAO().CreateUser(inspector)
	defer config.Config.DB.Delete(inspector)
	var inspectors []models.User
	config.Config.DB.Model(&inspectors).Select()
	machine1 := &models.Machine { Name: "Machine1", Period: 7, InspectorId: inspectors[0].Id }
	machineDao := daos.NewMachineDAO()
	machineDao.CreateMachine(machine1)
	defer config.Config.DB.Delete(machine1)
	machine2 := &models.Machine { Name: "Machine2", Period: 7, InspectorId: inspectors[0].Id, NextInspection: time.Now() }
	machineDao.CreateMachine(machine2)
	defer config.Config.DB.Delete(machine2)
	machine3 := &models.Machine { Name: "Machine3", Period: 7, InspectorId: inspectors[0].Id, NextInspection: time.Now().Add(-48*time.Hour) }
	machineDao.CreateMachine(machine3)
	defer config.Config.DB.Delete(machine3)
	machine4 := &models.Machine { Name: "Machine4", Period: 7, InspectorId: inspectors[0].Id, NextInspection: time.Now().Add(48*time.Hour) }
	machineDao.CreateMachine(machine4)
	defer config.Config.DB.Delete(machine4)
	machine5 := &models.Machine { Name: "Machine5", Period: 7 }
	machineDao.CreateMachine(machine5)
	defer config.Config.DB.Delete(machine5)
	machinesToInspect, err := machineDao.GetMachinesByInspectorIdAndRevision(inspectors[0].Id)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(machinesToInspect))
	assert.Equal(t, machine3.Name, machinesToInspect[0].Name)
	assert.Equal(t, machine2.Name, machinesToInspect[1].Name)
	assert.Equal(t, machine1.Name, machinesToInspect[2].Name)
}

func TestCreateMachine(t *testing.T) {
	company := &models.Company { Name: "company", Code: "4bb1500c-2b05-475c-bff2-ffe93942c697" }
	daos.NewCompanyDAO().CreateCompany(company)
	defer config.Config.DB.Delete(company)
	var companies []models.Company
	config.Config.DB.Model(&companies).Select()
	inspector := &models.User { Name: "testuser", IsAdmin: false, CompanyId: companies[0].Id }
	daos.NewUserDAO().CreateUser(inspector)
	defer config.Config.DB.Delete(inspector)
	var inspectors[] models.User
	config.Config.DB.Model(&inspectors).Select()
	machine := &models.Machine { InspectorId: inspectors[0].Id, Name: "testmachine", Period: 7, NextInspection: time.Now(), LastInspection: time.Now().Add(-24*7*time.Hour) }
	dao := daos.NewMachineDAO()
	err := dao.CreateMachine(machine)
	assert.Nil(t, err)
	defer config.Config.DB.Delete(machine)
	var machines []models.Machine
	config.Config.DB.Model(&machines).Relation("Inspector").Select()
	assert.Equal(t, 1, len(machines))
	assert.Equal(t, machine.Name, machines[0].Name)
	assert.Equal(t, machine.Period, machines[0].Period)
	assert.Equal(t, &inspectors[0], machines[0].Inspector)
	assert.True(t, datesEqual(machine.NextInspection, machines[0].NextInspection))
	assert.True(t, datesEqual(machine.LastInspection, machines[0].LastInspection))
}

func TestUpdateMachine(t *testing.T) {
	machine := &models.Machine { Name: "testmachine", Period: 7, NextInspection: time.Now(), LastInspection: time.Now().Add(-24*7*time.Hour) }
	dao := daos.NewMachineDAO()
	err := dao.CreateMachine(machine)
	assert.Nil(t, err)
	defer config.Config.DB.Delete(machine)
	machine.Name = "othername"
	machine.Period = 8
	err = dao.UpdateMachine(machine)
	assert.Nil(t, err)
	var machines []models.Machine
	config.Config.DB.Model(&machines).Select()
	assert.Equal(t, "othername", machines[0].Name)
	assert.Equal(t, 8, machines[0].Period)
}

func TestDeleteMachine(t *testing.T) {
	machine := &models.Machine { Name: "testmachine", Period: 7, NextInspection: time.Now(), LastInspection: time.Now().Add(-24*7*time.Hour) }
	dao := daos.NewMachineDAO()
	err := dao.CreateMachine(machine)
	assert.Nil(t, err)
	var machines []models.Machine
	config.Config.DB.Model(&machines).Select()
	assert.Equal(t, 1, len(machines))
	err = dao.DeleteMachine(machines[0].Id)
	assert.Nil(t, err)
	config.Config.DB.Model(&machines).Select()
	assert.Equal(t, 0, len(machines))
}

func TestGetMachineInspectionsByMachineId(t *testing.T) {
	company := &models.Company { Name: "company", Code: "4bb1500c-2b05-475c-bff2-ffe93942c697" }
	daos.NewCompanyDAO().CreateCompany(company)
	defer config.Config.DB.Delete(company)
	var companies []models.Company
	config.Config.DB.Model(&companies).Select()
	inspector := &models.User { Name: "testuser", IsAdmin: false, CompanyId: companies[0].Id }
	daos.NewUserDAO().CreateUser(inspector)
	defer config.Config.DB.Delete(inspector)
	machine := &models.Machine { Name: "testmachine", Period: 7, NextInspection: time.Now(), LastInspection: time.Now().Add(-24*7*time.Hour) }
	daos.NewMachineDAO().CreateMachine(machine)
	defer config.Config.DB.Delete(machine)
	var tempMachines[] models.Machine
	config.Config.DB.Model(&tempMachines).Select()
	machineInspection := &models.MachineInspection { MachineId: tempMachines[0].Id, Inspector: inspector, Date: time.Now(), Observations: "test observations" }
	dao := daos.NewMachineInspectionDAO()
	err := dao.CreateMachineInspection(machineInspection)
	assert.Nil(t, err)
	defer config.Config.DB.Delete(machineInspection)
	secondMachineInspection := &models.MachineInspection { MachineId: tempMachines[0].Id, Inspector: inspector, Date: time.Now().Add(24*7*time.Hour), Observations: "" }
	err = dao.CreateMachineInspection(secondMachineInspection)
	assert.Nil(t, err)
	defer config.Config.DB.Delete(secondMachineInspection)
	var inspections []models.MachineInspection
	inspections, err = dao.GetMachineInspectionsByMachineId(tempMachines[0].Id, false)
	assert.Equal(t, 2, len(inspections))
	assert.True(t, datesEqual(secondMachineInspection.Date, inspections[0].Date))
	assert.Equal(t, secondMachineInspection.Observations, inspections[0].Observations)
	assert.True(t, datesEqual(machineInspection.Date, inspections[1].Date))
	assert.Equal(t, machineInspection.Observations, inspections[1].Observations)
}

func TestGetMachineInspectionsByMachineIdOnlyObservations(t *testing.T) {
	company := &models.Company { Name: "company", Code: "4bb1500c-2b05-475c-bff2-ffe93942c697" }
	daos.NewCompanyDAO().CreateCompany(company)
	defer config.Config.DB.Delete(company)
	var companies []models.Company
	config.Config.DB.Model(&companies).Select()
	inspector := &models.User { Name: "testuser", IsAdmin: false, CompanyId: companies[0].Id }
	daos.NewUserDAO().CreateUser(inspector)
	defer config.Config.DB.Delete(inspector)
	machine := &models.Machine { Name: "testmachine", Period: 7, NextInspection: time.Now(), LastInspection: time.Now().Add(-24*7*time.Hour) }
	daos.NewMachineDAO().CreateMachine(machine)
	defer config.Config.DB.Delete(machine)
	var tempMachines[] models.Machine
	config.Config.DB.Model(&tempMachines).Select()
	machineInspection := &models.MachineInspection { MachineId: tempMachines[0].Id, Inspector: inspector, Date: time.Now(), Observations: "test observations" }
	dao := daos.NewMachineInspectionDAO()
	err := dao.CreateMachineInspection(machineInspection)
	assert.Nil(t, err)
	defer config.Config.DB.Delete(machineInspection)
	secondMachineInspection := &models.MachineInspection { MachineId: tempMachines[0].Id, Inspector: inspector, Date: time.Now().Add(24*7*time.Hour), Observations: "" }
	err = dao.CreateMachineInspection(secondMachineInspection)
	assert.Nil(t, err)
	defer config.Config.DB.Delete(secondMachineInspection)
	var inspections []models.MachineInspection
	inspections, err = dao.GetMachineInspectionsByMachineId(tempMachines[0].Id, true)
	assert.Equal(t, 1, len(inspections))
	assert.True(t, datesEqual(machineInspection.Date, inspections[0].Date))
	assert.Equal(t, machineInspection.Observations, inspections[0].Observations)
}

func TestCreateMachineInspection(t *testing.T) {
	company := &models.Company { Name: "company", Code: "4bb1500c-2b05-475c-bff2-ffe93942c697" }
	daos.NewCompanyDAO().CreateCompany(company)
	defer config.Config.DB.Delete(company)
	var companies []models.Company
	config.Config.DB.Model(&companies).Select()
	inspector := &models.User { Name: "testuser", IsAdmin: false, CompanyId: companies[0].Id }
	daos.NewUserDAO().CreateUser(inspector)
	defer config.Config.DB.Delete(inspector)
	var inspectors []models.User
	config.Config.DB.Model(&inspectors).Select()
	machine := &models.Machine { Name: "testmachine", Period: 7, NextInspection: time.Now(), LastInspection: time.Now().Add(-24*7*time.Hour) }
	daos.NewMachineDAO().CreateMachine(machine)
	defer config.Config.DB.Delete(machine)
	var tempMachines[] models.Machine
	config.Config.DB.Model(&tempMachines).Select()
	machineInspection := &models.MachineInspection { MachineId: tempMachines[0].Id, InspectorId: inspectors[0].Id, Date: time.Now(), Observations: "test observations" }
	dao := daos.NewMachineInspectionDAO()
	err := dao.CreateMachineInspection(machineInspection)
	assert.Nil(t, err)
	defer config.Config.DB.Delete(machineInspection)
	var machineInspections []models.MachineInspection
	config.Config.DB.Model(&machineInspections).Relation("Inspector").Select()
	assert.Equal(t, 1, len(machineInspections))
	assert.True(t, datesEqual(machineInspection.Date, machineInspections[0].Date))
	assert.Equal(t, machineInspection.Observations, machineInspections[0].Observations)
	assert.Equal(t, tempMachines[0].Id, machineInspections[0].MachineId)
	assert.Equal(t, inspectors[0].Id, machineInspections[0].Inspector.Id)
	var machines []models.Machine
	config.Config.DB.Model(&machines).Relation("Inspections").Relation("Inspections.Inspector").Select()
	assert.Equal(t, 1, len(machines))
	assert.Equal(t, tempMachines[0].Id, machines[0].Id)
	assert.Equal(t, 1, len(machines[0].Inspections))
	assert.True(t, datesEqual(machineInspection.Date, machines[0].Inspections[0].Date))
	assert.Equal(t, machineInspection.Observations, machines[0].Inspections[0].Observations)
	assert.Equal(t, tempMachines[0].Id, machines[0].Inspections[0].MachineId)
	assert.Equal(t, inspector.Id, machines[0].Inspections[0].Inspector.Id)
}

func TestGetUsersByCompanyId(t *testing.T) {
	company := &models.Company { Name: "company", Code: "4bb1500c-2b05-475c-bff2-ffe93942c697" }
	daos.NewCompanyDAO().CreateCompany(company)
	defer config.Config.DB.Delete(company)
	var companies []models.Company
	config.Config.DB.Model(&companies).Select()
	user := &models.User { Name: "testuser", IsAdmin: true, CompanyId: companies[0].Id }
	dao := daos.NewUserDAO()
	dao.CreateUser(user)
	defer config.Config.DB.Delete(user)
	users, err := dao.GetUsersByCompanyId(companies[0].Id)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, user.Name, users[0].Name)
	assert.Equal(t, user.IsAdmin, users[0].IsAdmin)
	users, err = dao.GetUsersByCompanyId(companies[0].Id + 1)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(users))
}

func TestCreateUser(t *testing.T) {
	company := &models.Company { Name: "company", Code: "4bb1500c-2b05-475c-bff2-ffe93942c697" }
	daos.NewCompanyDAO().CreateCompany(company)
	defer config.Config.DB.Delete(company)
	var companies []models.Company
	config.Config.DB.Model(&companies).Select()
	user := &models.User { Name: "testuser", IsAdmin: true, CompanyId: companies[0].Id }
	dao := daos.NewUserDAO()
	err := dao.CreateUser(user)
	assert.Nil(t, err)
	defer config.Config.DB.Delete(user)
	var users []models.User
	config.Config.DB.Model(&users).Select()
	assert.Equal(t, 1, len(users))
	assert.Equal(t, user.Name, users[0].Name)
	assert.Equal(t, user.IsAdmin, users[0].IsAdmin)
}

func TestUpdateUser(t *testing.T) {
	company := &models.Company { Name: "company", Code: "4bb1500c-2b05-475c-bff2-ffe93942c697" }
	daos.NewCompanyDAO().CreateCompany(company)
	defer config.Config.DB.Delete(company)
	var companies []models.Company
	config.Config.DB.Model(&companies).Select()
	user := &models.User { Name: "testuser", IsAdmin: true, CompanyId: companies[0].Id }
	dao := daos.NewUserDAO()
	dao.CreateUser(user)
	defer config.Config.DB.Delete(user)
	user.Name = "otheruser"
	user.IsAdmin = false
	err := dao.UpdateUser(user)
	assert.Nil(t, err)
	var users []models.User
	config.Config.DB.Model(&users).Select()
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "otheruser", users[0].Name)
	assert.Equal(t, false, users[0].IsAdmin)
}

func TestDeleteUser(t *testing.T) {
	company := &models.Company { Name: "company", Code: "4bb1500c-2b05-475c-bff2-ffe93942c697" }
	daos.NewCompanyDAO().CreateCompany(company)
	defer config.Config.DB.Delete(company)
	var companies []models.Company
	config.Config.DB.Model(&companies).Select()
	user := &models.User { Name: "testuser", IsAdmin: true, CompanyId: companies[0].Id }
	dao := daos.NewUserDAO()
	err := dao.CreateUser(user)
	assert.Nil(t, err)
	var users []models.User
	config.Config.DB.Model(&users).Select()
	assert.Equal(t, 1, len(users))
	err = dao.DeleteUser(users[0].Id)
	assert.Nil(t, err)
	config.Config.DB.Model(&users).Select()
	assert.Equal(t, 0, len(users))
}

func datesEqual(expected time.Time, actual time.Time) bool {
	return expected.Year() == actual.Year() &&
		expected.Month() == actual.Month() &&
		expected.Day() == actual.Day()
}