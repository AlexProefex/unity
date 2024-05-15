package types

type RecompensaRegister struct {
	Insignia uint `json:"insignia" binding:"required"`
	Usuario  uint `json:"usuario" binding:"required"`
}

type RecompensaUpdate struct {
	ID       uint `json:"id" binding:"required"`
	Insignia uint `json:"insignia" binding:"required"`
	Usuario  uint `json:"usuario" binding:"required"`
}
