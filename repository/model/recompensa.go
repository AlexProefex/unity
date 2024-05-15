package model

import "gorm.io/gorm"

type Recompensa struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	InsigniaId uint `gorm:"size:255;not null;" json:"insignia"`
	UsuarioId  uint
}
