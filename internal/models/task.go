package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UUID        string `gorm:"unique;not null;column:UUID"`
	Name        string `gorm:"size:255;not null;column:NAME"`
	Description string `gorm:"size:2047;not null;column:DESCRIPTION"`
	Status      string `gorm:"size:20;not null;column:STATUS"`
	CreatedBy   uint   `gorm:"not null;column:CREATED_BY"`
	UpdatedBy   uint   `gorm:"not null;column:UPDATED_BY"`
	DeletedBy   uint   `gorm:"column:DELETED_BY"`
}
