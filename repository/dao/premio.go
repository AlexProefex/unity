package dao

import (
	"errors"
	"unity/initialize"
	"unity/repository/model"
)

type Premio model.Premio

func GetAllPremios() ([]Premio, error) {

	var premio []Premio

	if err := initialize.DB.Find(&premio).Error; err != nil {
		return premio, errors.New("data not found")
	}

	return premio, nil
}

func (premio *Premio) SavePremio() (*Premio, error) {

	err := initialize.DB.Create(premio).Error

	if err != nil {
		return &Premio{}, errors.New("data not found")
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