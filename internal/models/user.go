package models

type User struct {
	Id int
	Name string
	IsAdmin bool
	Company *Company
}