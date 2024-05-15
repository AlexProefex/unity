package types

type LocacionRegister struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	Latitud     string `json:"latitud" binding:"required"`
	Longintud   string `json:"longintud" binding:"required"`
	Categoria   uint   `json:"categoria" binding:"required"`
}

type LocacionUpdate struct {
	ID          uint   `json:"id" binding:"required"`
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
	Latitud     string `json:"latitud" binding:"required"`
	Longintud   string `json:"longintud" binding:"required"`
	Categoria   uint   `json:"categoria" binding:"required"`
}
