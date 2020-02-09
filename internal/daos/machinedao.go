package daos

import (
	"github.com/fercarcedo/quadrant-server/internal/config"
	"github.com/fercarcedo/quadrant-server/internal/models"
)

type MachineDAO interface {
	GetMachinesByInspectorIdAndRevision(inspectorId int) ([]models.Machine, error)
	CreateMachine(machine *models.Machine) error
	UpdateMachine(machine *models.Machine) error
	DeleteMachine(machineId int) error
}

type machineDAO struct {}

func NewMachineDAO() MachineDAO {
	return &machineDAO{}
}

func (dao *machineDAO) GetMachinesByInspectorIdAndRevision(inspectorId int) ([]models.Machine, error) {
	var machines []models.Machine
	err := config.Config.DB.Model(&machines).
		Where("inspector_id = ?", inspectorId).
		Where("NOW() >= next_inspection OR next_inspection IS NULL").
		Order("next_inspection ASC").
		Select()
	return machines, err
}

func (dao *machineDAO) CreateMachine(machine *models.Machine) error {
	return config.Config.DB.Insert(machine)
}

func (dao *machineDAO) UpdateMachine(machine *models.Machine) error {
	return config.Config.DB.Update(machine)
}

func (dao *machineDAO) DeleteMachine(machineId int) error {
	machine := &models.Machine{ Id: machineId }
	return config.Config.DB.Delete(machine)
}