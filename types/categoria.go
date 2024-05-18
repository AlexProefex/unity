package types

import "time"

type CategoriaRegister struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	Insignia    string `json:"insignia" binding:"required"`
	Tiempo      string `json:"tiempo" binding:"required"`
} //@name Nueva Categoria

type CategoriaUpdate struct {
	ID          uint   `json:"id" binding:"required"`
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	Insignia    string `json:"insignia" binding:"required"`
	Tiempo      string `json:"tiempo" binding:"required"`
} //@name Actualizar Categoria

type CategoriaChallenge struct {
	ID      uint      `json:"id" binding:"required"`
	Nombre  string    `json:"nombre" binding:"required"`
	Usuario uint      `json:"usuario" binding:"required"`
	Tiempo  time.Time `json:"tiempo" binding:"required"`
} //@name Challenge
