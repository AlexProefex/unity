package dao

import (
	"errors"
	"unity/initialize"
	"unity/repository/model"
	"unity/utils"

	"gorm.io/gorm"
)

type Locacion model.Locacion

func GetAllLocacion() ([]Locacion, error) {
	var locacion []Locacion
	if err := initialize.DB.Find(&locacion).Error; err != nil {
		return locacion, errors.New(utils.Not_found)
	}
	return locacion, nil
}

func (locacion *Locacion) SaveLocacion() (*Locacion, error) {
	err := initialize.DB.Create(locacion).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Locacion{}, errors.New(utils.Not_found)
		}
		return &Locacion{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	return locacion, err
}

func (locacion *Locacion) UpdateLocacion(uid uint) (*Locacion, error) {
	err := initialize.DB.Save(locacion).Where("ID=?", uid).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Locacion{}, errors.New(utils.Not_found)
		}
		return &Locacion{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	return locacion, err
}

func (locacion Locacion) GetLocacionByID(uid uint) (Locacion, error) {

	if err := initialize.DB.First(&locacion, uid).Error; err != nil {
		return locacion, errors.New("user not found")
	}
	return locacion, nil
}

func GetLocacionInEvent(elements []uint) ([]Locacion, error) {
	var locacion []Locacion
	if err := initialize.DB.Debug().Where("id IN ?", elements).Find(&locacion).Error; err != nil {
		return locacion, errors.New(utils.Not_found)
	}
	return locacion, nil
}
