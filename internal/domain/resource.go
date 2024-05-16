package domain

import (
	"gorm.io/gorm"
)

type BurjoGroup struct {
	gorm.Model
	Name        string
	Owner       string
	Description string
	Burjos      []Burjo
}

type Burjo struct {
	gorm.Model
	Name         string
	Address      string
	Employees    []Employee
	BurjoGroupID uint
}

type Employee struct {
	gorm.Model
	Name    string
	Origin  string
	BurjoID uint
}
