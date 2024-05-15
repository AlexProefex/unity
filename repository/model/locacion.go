package model

import "gorm.io/gorm"

type Locacion struct {
	gorm.Model
	ID          uint              `gorm:"primaryKey"`
	Nombre      string            `gorm:"size:255;not null;" json:"nombre"`
	Descripcion string            `gorm:"size:255;not null;" json:"descripcion"`
	Latitud     string            `gorm:"size:255;not null;" json:"latitud"`
	Longintud   string            `gorm:"size:255;not null;" json:"longintud"`
	Usuario     []UsuarioLocacion `gorm:"foreignKey:LocacionId;"`
	CategoriaId uint
}
