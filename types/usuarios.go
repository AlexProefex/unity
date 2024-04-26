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
