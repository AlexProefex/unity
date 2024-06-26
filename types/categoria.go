package types

type CategoriaRegister struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	Insignia    string `json:"insignia" binding:"required"`
	Tiempo      int    `json:"tiempo" binding:"required"`
}

type CategoriaUpdate struct {
	ID          uint   `json:"id" binding:"required"`
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	Insignia    string `json:"insignia" binding:"required"`
	Tiempo      int    `json:"tiempo" binding:"required"`
}

type CategoriaChallenge struct {
	ID      uint   `json:"id" binding:"required"`
	Nombre  string `json:"nombre" binding:"required"`
	Usuario uint   `json:"usuario" binding:"required"`
}

type CategoriaMiniChallenge struct {
	Nombre  string `json:"nombre" binding:"required"`
	Usuario uint   `json:"usuario" binding:"required"`
}

type CategoriaRestar struct {
	KeyServer string `json:"keyserver" binding:"required"`
}
