package types

import "time"

type UsuariosRegister struct {
	Nombre             string    `json:"nombre" binding:"required"`
	Apellidos          string    `json:"apellidos" binding:"required"`
	Correo_electronico string    `json:"Correo_electronico" binding:"required,email"`
	Codigo_pais        string    `json:"codigo_pais" binding:"required"`
	Celular            string    `json:"celular" binding:"required"`
	Genero             string    `json:"genero" binding:"required"`
	Fecha_nacimiento   time.Time `json:"fecha_nacimiento" binding:"required"`
	Nacionalidad       string    `json:"nacionalidad" binding:"required"`
	Password           string    `json:"password" binding:"required"`
	Secret             string    `json:"secret" binding:"required"`
} //@name Registro Usuario

type UsuariosLogin struct {
	Correo_electronico string `json:"Correo_electronico" binding:"required"`
	Password           string `json:"password" binding:"required"`
} //@name Login

type UsuariosModel struct {
	Nombre             string    `json:"nombre" binding:"required"`
	Apellidos          string    `json:"apellidos" binding:"required"`
	Correo_electronico string    `json:"Correo_electronico" binding:"required,email"`
	Codigo_pais        string    `json:"codigo_pais" binding:"required"`
	Celular            string    `json:"celular" binding:"required"`
	Genero             string    `json:"genero" binding:"required"`
	Fecha_nacimiento   time.Time `json:"fecha_nacimiento" binding:"required"`
	Nacionalidad       string    `json:"nacionalidad" binding:"required"`
} //@name Usuario

type UsuariosInsignia struct {
	Cantidad int `json:"cantidad" binding:"required"`
} //@name Insignia

type UsuariosPuntos struct {
	Puntos int `json:"puntos" binding:"required"`
} //@name Puntos

type UsuariosValidateCP struct {
	Producto int `json:"producto" binding:"required"`
	Cantidad int `json:"cantidad"`
	Puntos   int `json:"puntos" `
} //@name CanjearPuntos

type UsuariosPassword struct {
	Correo_electronico string `json:"Correo_electronico" binding:"required,email"`
	Password           string `json:"password" binding:"required"`
	Secret             string `json:"secret" binding:"required"`
} //@name ValidarPassword

type ConfirmPassword struct {
	Correo_electronico string `json:"Correo_electronico" binding:"required,email"`
	Password           string `json:"password" binding:"required"`
	NewPassword        string `json:"newpassword" binding:"required"`
} //@name CambiarPassword

type UpdatePerfil struct {
	Nombre    string `json:"nombre" binding:"required"`
	Apellidos string `json:"apellidos" binding:"required"`
} //@name Perfil

type AsingarPuntos struct {
	ID        uint   `json:"id" binding:"required"`
	Puntos    int    `json:"puntos" binding:"required"`
	Insignia  uint   `json:"insignia" binding:"required"`
	TypeRoute string `json:"type" binding:"required"`
} //@name Reclamar Puntos
