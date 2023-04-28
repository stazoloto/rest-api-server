package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid"`
	Title    string
	SubTitle string
	Text     string
}
