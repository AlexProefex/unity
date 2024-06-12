package types

type PremioRegister struct {
	Nombre         string  `json:"nombre" binding:"required"`
	Descripcion    string  `json:"descripcion" binding:"required"`
	Imagen         string  `json:"imagen" binding:"required"`
	Tipo           string  `json:"Tipo" binding:"required"`
	Descuento      float32 `json:"descuento" binding:"required"`
	MontoDescuento float32 `json:"monto" binding:"required"`
}

type PremioUpdate struct {
	ID             uint    `json:"id" binding:"required"`
	Nombre         string  `json:"nombre" binding:"required"`
	Descripcion    string  `json:"descripcion" binding:"required"`
	Imagen         string  `json:"imagen" binding:"required"`
	Tipo           string  `json:"Tipo" binding:"required"`
	Descuento      float32 `json:"descuento" binding:"required"`
	MontoDescuento float32 `json:"monto" binding:"required"`
}
