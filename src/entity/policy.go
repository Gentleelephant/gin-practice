package entity

import "gorm.io/gorm"

type Policy struct {
	gorm.Model

	PType string `gorm:"column:ptype"`

	Object string `gorm:"column:object"`

	Resource string `gorm:"column:resource"`

	Operation string `gorm:"column:operation"`
}
