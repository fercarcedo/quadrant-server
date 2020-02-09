package models

import "time"

type Machine struct {
	Id int
	Name string
	Period int
	InspectorId int
	Inspector *User
	NextInspection time.Time `pg:"type:date"`
	LastInspection time.Time `pg:"type:date"`
	Inspections []*MachineInspection
}