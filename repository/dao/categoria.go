package dao

import (
	"errors"
	"unity/initialize"
	"unity/repository/model"
	"unity/utils"

	"gorm.io/gorm"
)

type Categoria model.Categoria

func GetAllCategorias() ([]Categoria, error) {
	var categoria []Categoria
	if err := initialize.DB.Find(&categoria).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return categoria, errors.New(utils.Not_found)
		}
		return categoria, errors.New(utils.Ha_ocurrido_un_error)
	}
	return categoria, nil
}

func (categoria *Categoria) SaveCategorias() (*Categoria, error) {
	err := initialize.DB.Create(categoria).Error
	if err != nil {
		return &Categoria{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	return categoria, err
}

func (categoria *Categoria) UpdateCategorias(uid uint) (*Categoria, error) {
	err := initialize.DB.Save(categoria).Where("ID=?", uid).Error
	if err != nil {
		return &Categoria{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	return categoria, err
}

func (categoria Categoria) GetCategoriaByID(uid uint) (Categoria, error) {
	if err := initialize.DB.First(&categoria, uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return categoria, errors.New(utils.Not_found)
		}
		return categoria, errors.New(utils.Ha_ocurrido_un_error)
	}
	return categoria, nil
}

func GetChallenge(uid uint) ([]Categoria, error) {
	var categoria []Categoria
	if err := initialize.DB.Where(&Categoria{ID: uid}).Model(&Categoria{}).Preload("Locaciones").Find(&categoria).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return categoria, errors.New(utils.Not_found)
		}
		return categoria, errors.New(utils.Ha_ocurrido_un_error)
	}
	return categoria, nil
}
