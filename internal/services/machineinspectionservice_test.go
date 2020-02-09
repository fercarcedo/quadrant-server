package services_test

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/fercarcedo/quadrant-server/internal/models"
	"github.com/fercarcedo/quadrant-server/internal/services"
)

type mockMachineInspectionDAO struct {
	mock.Mock
}

func (m *mockMachineInspectionDAO) GetMachineInspectionsByMachineId(machineId int, onlyObservations bool) ([]models.MachineInspection, error) {
	args := m.Called(machineId, onlyObservations)
	return args.Get(0).([]models.MachineInspection), args.Error(1)
}

func (m *mockMachineInspectionDAO) CreateMachineInspection(machineInspection *models.MachineInspection) error {
	args := m.Called(machineInspection)
	return args.Error(0)
}

func TestGetMachineInspectionsByMachineId(t *testing.T) {
	dao := new(mockMachineInspectionDAO)
	machineInspections := []models.MachineInspection{ 
		models.MachineInspection{ Id: 1, MachineId: 2, Date: time.Now(), Observations: "obs1"}, 
		models.MachineInspection{ Id: 2, MachineId: 3, Date: time.Now(), Observations: "obs2"}, 
	}
	dao.On("GetMachineInspectionsByMachineId", 1, false).Return(machineInspections, nil)
	service := services.NewMachineInspectionService(dao)
	inspections, err := service.GetMachineInspectionsByMachineId(1, false)
	assert.Nil(t, err)
	assert.Equal(t, machineInspections, inspections)
	dao.AssertExpectations(t)
}

func TestCreateMachineInspection(t *testing.T) {
	dao := new(mockMachineInspectionDAO)
	machineInspection := &models.MachineInspection {
		MachineId: 2,
		Date: time.Now(),
		Observations: "obs1",
	}
	dao.On("CreateMachineInspection", machineInspection).Return(nil)
	service := services.NewMachineInspectionService(dao)
	assert.Nil(t, service.CreateMachineInspection(machineInspection))
	dao.AssertExpectations(t)
}