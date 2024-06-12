package service

import (
	"errors"
	"unity/repository/dao"
	"unity/types"
)

func ServiceGetAllPremio() ([]dao.Premio, error) {
	premio, err := dao.GetAllPremios()
	return premio, err
}

func ServiceGetAllPremioRegalo() ([]dao.Premio, error) {
	premio, err := dao.GetAllPremiosRegalos()
	return premio, err
}

func ServiceGetAllPremioDescuento() ([]dao.Premio, error) {
	premio, err := dao.GetAllPremiosDescuentos()
	return premio, err
}

func ServiceSavePremio(input types.PremioRegister) (dao.Premio, error) {
	premio := dao.Premio{
		Nombre:         input.Nombre,
		Descripcion:    input.Descripcion,
		Imagen:         input.Imagen,
		Tipo:           input.Tipo,
		Descuento:      input.Descuento,
		MontoDescuento: input.MontoDescuento,
	}
	current, err := premio.SavePremio()
	return *current, err
}

func ServiceUpdatePremio(input types.PremioUpdate, id uint) (dao.Premio, error) {

	premio := dao.Premio{
		ID:             input.ID,
		Nombre:         input.Nombre,
		Descripcion:    input.Descripcion,
		Imagen:         input.Imagen,
		Tipo:           input.Tipo,
		Descuento:      input.Descuento,
		MontoDescuento: input.MontoDescuento,
	}

	if premio.ID != id {
		return premio, errors.New("no se pudo actualizar el recurso solicitado")
	}
	current, err := premio.UpdatePremio(id)
	return *current, err
}

func ServiceGetPremioByID(uid uint) (dao.Premio, error) {
	premio := dao.Premio{
		ID: uid,
	}
	Premio, err := premio.GetPremioByID(uid)
	return Premio, err
}
