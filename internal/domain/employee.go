package domain

import (
	"fmt"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name    string
	Origin  string
	BurjoID uint
}

func (e *Employee) BeforeUpdate(_ *gorm.DB) error {
	if e.ID == 0 {
		return fmt.Errorf("no id provided")
	}

	return nil
}

func (e *Employee) BeforeDelete(_ *gorm.DB) error {
	if e.ID == 0 {
		return fmt.Errorf("no id provided")
	}

	return nil
}
