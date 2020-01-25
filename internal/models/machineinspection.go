package models

import "time"

type MachineInspector struct {
	Id int
	Machine *Machine
	Inspector *User
	Date time.Time
	Observations string
}