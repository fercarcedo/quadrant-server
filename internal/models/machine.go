package models

import "time"

type Machine struct {
	Id int `json:"-"`
	Name string `pg:",notnull" json:"name"`
	Period int `pg:",notnull" json:"period"`
	InspectorId int `json:"inspector_id"`
	Inspector *User `json:"-"`
	NextInspection time.Time `pg:"type:date" json:"next_inspection"`
	LastInspection time.Time `pg:"type:date" json:"last_inspection"`
	Inspections []*MachineInspection `json:"-"`
}