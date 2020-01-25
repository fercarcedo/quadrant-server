package models

import "time"

type Machine struct {
	Id int
	Name string
	Period int
	Inspector *User
	NextInspection time.Time
	LastInspection time.Time
}