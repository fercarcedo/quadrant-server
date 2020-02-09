package services

import (
	"github.com/fercarcedo/quadrant-server/internal/daos"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type MachineService interface {
	GetMachinesByInspectorIdAndRevision(inspectorId int) ([]models.Machine, error)
	CreateMachine(machine *models.Machine) error
	UpdateMachine(machine *models.Machine) error
	DeleteMachine(machineId int) error
}

type machineService struct {
	dao daos.MachineDAO
}

func NewMachineService(dao daos.MachineDAO) MachineService {
	return &machineService{ dao }
}

func (service *machineService) GetMachinesByInspectorIdAndRevision(inspectorId int) ([]models.Machine, error) {
	return service.dao.GetMachinesByInspectorIdAndRevision(inspectorId)
}

func (service *machineService) CreateMachine(machine *models.Machine) error {
	return service.dao.CreateMachine(machine)
}

func (service *machineService) UpdateMachine(machine *models.Machine) error {
	return service.dao.UpdateMachine(machine)
}

func (service *machineService) DeleteMachine(machineId int) error {
	return service.dao.DeleteMachine(machineId)
}