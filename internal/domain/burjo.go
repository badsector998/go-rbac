package domain

import (
	"fmt"

	"gorm.io/gorm"
)

type Burjo struct {
	gorm.Model
	Name         string
	Address      string
	Employees    []Employee
	BurjoGroupID uint
}

func (b *Burjo) BeforeUpdate(_ *gorm.DB) error {
	if b.ID == 0 {
		return fmt.Errorf("no id provided")
	}

	return nil
}

func (b *Burjo) BeforeDelete(_ *gorm.DB) error {
	if b.ID == 0 {
		return fmt.Errorf("no id provided")
	}

	return nil
}
