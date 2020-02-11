package daos

import (
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type MachineInspectionDAO interface {
	GetMachineInspectionsByMachineId(machineId int, onlyObservations bool) ([]models.MachineInspection, error)
	CreateMachineInspection(machineInspection *models.MachineInspection) error
}

type machineInspectionDAO struct {}

func NewMachineInspectionDAO() MachineInspectionDAO {
	return &machineInspectionDAO{}
}

func (dao *machineInspectionDAO) GetMachineInspectionsByMachineId(machineId int, onlyObservations bool) ([]models.MachineInspection, error) {
	machineInspections := []models.MachineInspection{}
	query := config.Config.DB.Model(&machineInspections).Where("machine_id = ?", machineId)
	if (onlyObservations) {
		query = query.Where("(observations = '') IS FALSE")
	}
	err := query.Order("date DESC").Select()
	return machineInspections, err
}

func (dao *machineInspectionDAO) CreateMachineInspection(machineInspection *models.MachineInspection) error {
	return config.Config.DB.Insert(machineInspection)
}