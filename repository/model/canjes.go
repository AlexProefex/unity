package model

import (
	"gorm.io/gorm"
)

type Canjes struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	ProductId string `gorm:"size:255;not null;" json:"product_id"`
	UsuarioId string `gorm:"size:255;not null;" json:"usuario_id"`
	Puntos    string `gorm:"size:255;not null;" json:"puntos"`
	Insignias string `gorm:"size:255;not null;" json:"insignias"`
}
