package models

import (
	"gorm.io/gorm"
)

type Images struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null"`
	Path string `gorm:"type:varchar(100);not null"`
}
