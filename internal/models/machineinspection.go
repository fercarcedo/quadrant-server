package models

import "time"

type MachineInspection struct {
	Id int
	Inspector *User
	Date time.Time
	Observations string
}