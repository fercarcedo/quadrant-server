package models

type User struct {
	Id int `json:"-"`
	Name string `pg:",notnull" json:"name"`
	IsAdmin bool `json:"is_admin"`
	CompanyId int `pg:",notnull" json:"company_id"`
}