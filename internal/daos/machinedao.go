package daos

import (
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type MachineDAO struct {}

func NewMachineDAO() *MachineDAO {
	return &MachineDAO{}
}

func (dao *MachineDAO) GetMachinesByInspectorIdAndRevision(inspectorId int) ([]models.Machine, error) {
	var machines []models.Machine
	err := config.Config.DB.Model(&machines).
		Where("inspector_id = ?", inspectorId).
		Where("NOW() >= next_inspection OR next_inspection IS NULL").
		Order("next_inspection ASC").
		Select()
	return machines, err
}

func (dao *MachineDAO) CreateMachine(machine *models.Machine) error {
	return config.Config.DB.Insert(machine)
}

func (dao *MachineDAO) UpdateMachine(machine *models.Machine) error {
	return config.Config.DB.Update(machine)
}

func (dao *MachineDAO) DeleteMachine(machineId int) error {
	machine := &models.Machine{ Id: machineId }
	return config.Config.DB.Delete(machine)
}