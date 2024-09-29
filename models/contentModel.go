package models

import (
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	Title string `gorm:"type:varchar(100);not null"`
	Link string `gorm:"type:varchar(100);not null"`
	Desc string `gorm:"type:varchar(100);not null"`
}