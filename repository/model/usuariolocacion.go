package model

import (
	"gorm.io/gorm"
)

type UsuarioLocacion struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	UsuarioId  uint
	LocacionId uint
	Evento     string `gorm:"size:255;not null;" json:"evento"`
	Estado     string `gorm:"size:255;not null;" json:"estado"`
}
