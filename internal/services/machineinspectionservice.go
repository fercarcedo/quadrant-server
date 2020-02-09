package services

import (
	"github.com/fercarcedo/quadrant-server/internal/daos"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type MachineInspectionService interface {
	GetMachineInspectionsByMachineId(machineId int, onlyObservations bool) ([]models.MachineInspection, error)
	CreateMachineInspection(machineInspection *models.MachineInspection) error
}

type machineInspectionService struct {
	dao daos.MachineInspectionDAO
}

func NewMachineInspectionService(dao daos.MachineInspectionDAO) MachineInspectionService {
	return &machineInspectionService{ dao }
}

func (service *machineInspectionService) GetMachineInspectionsByMachineId(machineId int, onlyObservations bool) ([]models.MachineInspection, error) {
	return service.dao.GetMachineInspectionsByMachineId(machineId, onlyObservations)
}

func (service *machineInspectionService) CreateMachineInspection(machineInspection *models.MachineInspection) error {
	return service.dao.CreateMachineInspection(machineInspection)
}