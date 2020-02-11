package models

type Company struct {
	Id int `json:"-"`
	Name string `pg:",notnull" json:"name"`
	Code string `pg:"type:uuid,notnull,unique" json:"code"`
}