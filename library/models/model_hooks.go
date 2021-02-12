package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (oc *OffenseCategory) BeforeCreate(tx *gorm.DB) (err error) {
	oc.ID = uuid.New().String()
	return
}

func (oc *Human) BeforeCreate(tx *gorm.DB) (err error) {
	oc.ID = uuid.New().String()
	return
}

func (oc *Offense) BeforeCreate(tx *gorm.DB) (err error) {
	oc.ID = uuid.New().String()
	return
}
