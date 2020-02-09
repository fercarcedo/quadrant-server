package services_test

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/fercarcedo/quadrant-server/internal/models"
	"github.com/fercarcedo/quadrant-server/internal/services"
)

type mockMachineDAO struct {
	mock.Mock
}

func (m *mockMachineDAO) GetMachinesByInspectorIdAndRevision(inspectorId int) ([]models.Machine, error) {
	args := m.Called(inspectorId)
	return args.Get(0).([]models.Machine), args.Error(1)
}

func (m *mockMachineDAO) CreateMachine(machine *models.Machine) error {
	args := m.Called(machine)
	return args.Error(0)
}

func (m *mockMachineDAO) UpdateMachine(machine *models.Machine) error {
	args := m.Called(machine)
	return args.Error(0)
}

func (m *mockMachineDAO) DeleteMachine(machineId int) error {
	args := m.Called(machineId)
	return args.Error(0)
}

func TestGetMachinesByInspectorIdAndRevision(t *testing.T) {
	dao := new(mockMachineDAO)
	machines := []models.Machine{ 
		models.Machine{ Id: 1, Name: "machine1", Period: 5, InspectorId: 3, NextInspection: time.Now(), LastInspection: time.Now().Add(-24*7*time.Hour)}, 
		models.Machine{ Id: 2, Name: "machine2", Period: 7, InspectorId: 3, NextInspection: time.Now(), LastInspection: time.Now().Add(-24*7*time.Hour)}, 
	}
	dao.On("GetMachinesByInspectorIdAndRevision", 3).Return(machines, nil)
	service := services.NewMachineService(dao)
	machinesByInspectorId, err := service.GetMachinesByInspectorIdAndRevision(3)
	assert.Nil(t, err)
	assert.Equal(t, machines, machinesByInspectorId)
	dao.AssertExpectations(t)
}

func TestCreateMachine(t *testing.T) {
	dao := new(mockMachineDAO)
	machine := &models.Machine{
		Id: 1, 
		Name: "machine1", 
		Period: 5, 
		InspectorId: 3, 
		NextInspection: time.Now(),
		LastInspection: time.Now().Add(-24*7*time.Hour),
	} 
	dao.On("CreateMachine", machine).Return(nil)
	service := services.NewMachineService(dao)
	assert.Nil(t, service.CreateMachine(machine))
	dao.AssertExpectations(t)
}

func TestUpdateMachine(t *testing.T) {
	dao := new(mockMachineDAO)
	machine := &models.Machine{
		Id: 1, 
		Name: "machine1", 
		Period: 5, 
		InspectorId: 3, 
		NextInspection: time.Now(),
		LastInspection: time.Now().Add(-24*7*time.Hour),
	}
	dao.On("UpdateMachine", machine).Return(nil)
	service := services.NewMachineService(dao)
	assert.Nil(t, service.UpdateMachine(machine))
	dao.AssertExpectations(t)
}

func TestDeleteMachine(t *testing.T) {
	dao := new(mockMachineDAO)
	dao.On("DeleteMachine", 4).Return(nil)
	service := services.NewMachineService(dao)
	assert.Nil(t, service.DeleteMachine(4))
	dao.AssertExpectations(t)
}