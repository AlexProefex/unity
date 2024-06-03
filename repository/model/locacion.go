package model

import "gorm.io/gorm"

type Locacion struct {
	gorm.Model
	ID          uint              `gorm:"primaryKey"`
	Nombre      string            `gorm:"size:255;not null;" json:"nombre"`
	Descripcion string            `gorm:"size:255;not null;" json:"descripcion"`
	Latitud     string            `gorm:"size:255;not null;" json:"latitud"`
	Longitud    string            `gorm:"size:255;not null;" json:"longitud"`
	Usuario     []UsuarioLocacion `gorm:"foreignKey:LocacionId;"`
	CategoriaId uint
}
