package models

type Company struct {
	Id int
	Name string
	Code string `pg:"type:uuid,unique"`
}