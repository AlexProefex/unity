package types

type UsuariosRegister struct {
	Nombre             string `json:"nombre" binding:"required"`
	Apellidos          string `json:"apellidos" binding:"required"`
	Correo_electronico string `json:"Correo_electronico" binding:"required,email"`
	Codigo_pais        string `json:"codigo_pais" binding:"required"`
	Celular            string `json:"celular" binding:"required"`
	Genero             string `json:"genero" binding:"required"`
	Fecha_nacimiento   string `json:"fecha_nacimiento" binding:"required"`
	Nacionalidad       string `json:"nacionalidad" binding:"required"`
	Password           string `json:"password" binding:"required"`
}

type UsuariosLogin struct {
	Correo_electronico string `json:"Correo_electronico" binding:"required"`
	Password           string `json:"password" binding:"required"`
}

type UsuariosModel struct {
	Nombre             string `json:"nombre" binding:"required"`
	Apellidos          string `json:"apellidos" binding:"required"`
	Correo_electronico string `json:"Correo_electronico" binding:"required,email"`
	Codigo_pais        string `json:"codigo_pais" binding:"required"`
	Celular            string `json:"celular" binding:"required"`
	Genero             string `json:"genero" binding:"required"`
	Fecha_nacimiento   string `json:"fecha_nacimiento" binding:"required"`
	Nacionalidad       string `json:"nacionalidad" binding:"required"`
}

type UsuariosInsignia struct {
	Cantidad int `json:"cantidad" binding:"required"`
}

type UsuariosPuntos struct {
	Puntos int `json:"puntos" binding:"required"`
}

type UsuariosValidateCP struct {
	Cantidad int `json:"cantidad"`
	Puntos   int `json:"puntos"`
}

type UsuariosPassword struct {
	Correo_electronico string `json:"Correo_electronico" binding:"required,email"`
	Password           string `json:"password" binding:"required"`
	Secret             string `json:"secret" binding:"required"`
}

type ConfirmPassword struct {
	Correo_electronico string `json:"Correo_electronico" binding:"required,email"`
	Password           string `json:"password" binding:"required"`
	NewPassword        string `json:"newpassword" binding:"required"`
}
