package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name        string
	Adress      string
	Age         int
	CreditCards []CreditCard `gorm:"foreignKey:EmployeeId"`
}

type CreditCard struct {
	gorm.Model
	Number     string
	EmployeeId uint
}
