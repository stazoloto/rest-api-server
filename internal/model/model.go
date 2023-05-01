package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id     uuid.UUID `gorm:"type:uuid"`
	Title  string
	Author string
	Rating string
}
