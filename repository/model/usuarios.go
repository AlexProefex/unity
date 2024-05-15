package model

import (
	"gorm.io/gorm"
)

type Usuarios struct {
	gorm.Model
	ID                 uint              `gorm:"primaryKey"`
	Nombre             string            `gorm:"size:255;not null;" json:"nombre"`
	Apellidos          string            `gorm:"size:255;not null;" json:"apellidos"`
	Correo_electronico string            `gorm:"size:255;not null;uniqueIndex" json:"correo_electronico"`
	Codigo_pais        string            `gorm:"size:255;not null;" json:"codigo_pais"`
	Celular            string            `gorm:"size:255;not null;" json:"celular"`
	Genero             string            `gorm:"size:255;not null;" json:"genero"`
	Fecha_nacimiento   string            `gorm:"size:255;not null;" json:"fecha_nacimiento"`
	Nacionalidad       string            `gorm:"size:255;not null;" json:"nacionalidad"`
	Password           string            `gorm:"size:255;not null;" json:"-"`
	Puntos             int               `gorm:"default:0" json:"puntos"`
	Locacion           []UsuarioLocacion `gorm:"foreignKey:UsuarioId;"`
	Recompensa         []Recompensa      `gorm:"foreignKey:UsuarioId;"`
	Cantidad           int               `gorm:"default:0" json:"cantidad"`
	Secret             string            `gorm:"size:255;not null;" json:"secret"`
}
