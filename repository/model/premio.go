package model

import "gorm.io/gorm"

type Premio struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Nombre      string `gorm:"size:255;not null;" json:"insignia"`
	Descripcion string `gorm:"size:255;not null;" json:"descripcion"`
	Imagen      string `gorm:"size:255;not null;" json:"imagen"`
	Tipo        string `gorm:"size:255;not null;" json:"tipo"`
	Descuento   string `gorm:"size:255;not null;" json:"descuento"`
}