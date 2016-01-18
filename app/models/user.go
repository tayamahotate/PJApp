package models

type User struct {
	Id         int
	Username   string     `sql:"type:varchar(19);not null;unique"`
	Password   string     `sql:"type:varchar(255);not null"`
	Status     string     `sql:"type:char(1);not null"`
}
