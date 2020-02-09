package daos

import (
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type MachineInspectionDAO struct {}

func NewMachineInspectionDAO() *MachineInspectionDAO {
	return &MachineInspectionDAO{}
}

func (dao *MachineInspectionDAO) GetMachineInspectionsByMachineId(machineId int, onlyObservations bool) ([]models.MachineInspection, error) {
	var machineInspections []models.MachineInspection
	query := config.Config.DB.Model(&machineInspections).Where("machine_id = ?", machineId)
	if (onlyObservations) {
		query = query.Where("(observations = '') IS FALSE")
	}
	err := query.Order("date DESC").Select()
	return machineInspections, err
}

func (dao *MachineInspectionDAO) CreateMachineInspection(machineInspection *models.MachineInspection) error {
	return config.Config.DB.Insert(machineInspection)
}