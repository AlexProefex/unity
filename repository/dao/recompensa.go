package dao

import (
	"errors"
	"unity/initialize"
	"unity/repository/model"
)

type Recompensa model.Recompensa

func GetAllRecompensas() ([]Recompensa, error) {

	var recompensa []Recompensa

	if err := initialize.DB.Find(&recompensa).Error; err != nil {
		return recompensa, errors.New("data not found")
	}

	return recompensa, nil
}

func (recompensa *Recompensa) SaveRecompensa() (*Recompensa, error) {

	err := initialize.DB.Create(recompensa).Error

	if err != nil {
		return &Recompensa{}, errors.New("user not found")
	}
	return recompensa, err
}

func (recompensa *Recompensa) UpdateRecompensa(uid uint) (*Recompensa, error) {

	err := initialize.DB.Save(recompensa).Where("ID=?", uid).Error

	if err != nil {
		return &Recompensa{}, errors.New("user not found")
	}
	return recompensa, err
}

func (recompensa Recompensa) GetRecomenpsaByID(uid uint) (Recompensa, error) {

	if err := initialize.DB.First(&recompensa, uid).Error; err != nil {
		return recompensa, errors.New("user not found")
	}
	return recompensa, nil
}

func GetAllRecompensasByUserId(uid uint) ([]Recompensa, error) {

	var recompensa []Recompensa

	if err := initialize.DB.Where("usuario_id", uid).Distinct("insignia_id").Find(&recompensa).Error; err != nil {
		return recompensa, err
	}

	return recompensa, nil
}
