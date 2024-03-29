package models

import (
	"gorm.io/gorm"
)

type Command struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Script string
}
