package model

import (
	"time"

	"gorm.io/gorm"
)

type Categoria struct {
	gorm.Model
	ID          uint       `gorm:"primaryKey"`
	Nombre      string     `gorm:"size:255;not null;" json:"nombre"`
	Descripcion string     `gorm:"size:255;not null;" json:"descripcion"`
	Insignia    string     `gorm:"size:255;not null;" json:"insignia"`
	Tiempo      time.Time  `gorm:"size:255;not null;" json:"tiempo"`
	Locaciones  []Locacion `gorm:"foreignKey:CategoriaId"`
}
