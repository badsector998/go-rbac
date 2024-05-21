package domain

import (
	"fmt"

	"gorm.io/gorm"
)

type BurjoGroup struct {
	gorm.Model
	Name        string
	Owner       string
	Description string
	Burjos      []Burjo
}

func (b *BurjoGroup) BeforeUpdate(_ *gorm.DB) error {
	if b.ID == 0 {
		return fmt.Errorf("no id provided")
	}

	return nil
}

func (b *BurjoGroup) BeforeDelete(_ *gorm.DB) error {
	if b.ID == 0 {
		return fmt.Errorf("no id provided")
	}

	return nil
}
