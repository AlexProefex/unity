package service

import (
	"errors"
	"unity/repository/dao"
	"unity/types"
	"unity/utils"
)

func ServiceGetAllLocacion() ([]dao.Locacion, error) {
	locacion, err := dao.GetAllLocacion()
	return locacion, err
}

func ServiceSaveLocacion(input types.LocacionRegister) (dao.Locacion, error) {
	locacion := dao.Locacion{
		Nombre:      input.Nombre,
		Descripcion: input.Descripcion,
		Latitud:     input.Latitud,
		Longintud:   input.Longintud,
		CategoriaId: input.Categoria,
	}
	current, err := locacion.SaveLocacion()
	return *current, err
}

func ServiceUpdateLocacion(input types.LocacionUpdate, id uint) (dao.Locacion, error) {
	locacion := dao.Locacion{
		ID:          input.ID,
		Nombre:      input.Nombre,
		Descripcion: input.Descripcion,
		Latitud:     input.Latitud,
		Longintud:   input.Longintud,
		CategoriaId: input.Categoria,
	}
	if locacion.ID != id {
		return locacion, errors.New(utils.InvalidID)
	}
	current, err := locacion.UpdateLocacion(id)
	return *current, err
}

func ServiceGetLocacionByID(uid uint) (dao.Locacion, error) {
	locacion := dao.Locacion{}
	locacion, err := locacion.GetLocacionByID(uid)
	return locacion, err
}
