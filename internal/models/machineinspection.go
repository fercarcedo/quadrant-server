package models

import "time"

type MachineInspection struct {
	Id int `json:"-"`
	MachineId int `pg:",notnull" json:"machine_id"`
	InspectorId int `json:"inspector_id"`
	Inspector *User `json:"-"`
	Date time.Time `pg:",notnull" json:"date"`
	Observations string `json:"observations"`
}