package dao

import (
	"errors"
	"unity/initialize"
	"unity/repository/model"
	"unity/utils"

	"gorm.io/gorm"
)

type Premio model.Premio

func GetAllPremios() ([]Premio, error) {
	var premio []Premio
	if err := initialize.DB.Find(&premio).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return premio, errors.New(utils.Not_found)
		}
		return premio, errors.New(utils.Ha_ocurrido_un_error)
	}
	return premio, nil
}

func GetAllPremiosRegalos() ([]Premio, error) {
	var premio []Premio
	if err := initialize.DB.Where("tipo =?", utils.PremioRegalo).Find(&premio).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return premio, errors.New(utils.Not_found)
		}
		return premio, errors.New(utils.Ha_ocurrido_un_error)
	}
	return premio, nil
}

func GetAllPremiosDescuentos() ([]Premio, error) {
	var premio []Premio
	if err := initialize.DB.Where("tipo=?", utils.PremioDescuento).Find(&premio).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return premio, errors.New(utils.Not_found)
		}
		return premio, errors.New(utils.Ha_ocurrido_un_error)
	}
	return premio, nil
}

func (premio *Premio) SavePremio() (*Premio, error) {
	err := initialize.DB.Create(premio).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return premio, errors.New(utils.Not_found)
		}
		return premio, errors.New(utils.Ha_ocurrido_un_error)
	}
	return premio, err
}

func (premio *Premio) UpdatePremio(uid uint) (*Premio, error) {

	err := initialize.DB.Save(premio).Where("ID=?", uid).Error

	if err != nil {
		return &Premio{}, errors.New("data not found")
	}
	return premio, err
}

func (premio Premio) GetPremioByID(uid uint) (Premio, error) {

	if err := initialize.DB.First(&premio, uid).Error; err != nil {
		return premio, errors.New("data not found")
	}
	return premio, nil
}
