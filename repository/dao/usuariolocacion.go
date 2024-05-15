package dao

import (
	"errors"
	"unity/initialize"
	"unity/repository/model"
)

type UsuarioLocacion model.UsuarioLocacion

func GetAllLocaionUsuario() ([]UsuarioLocacion, error) {

	var locacion_usuario []UsuarioLocacion

	if err := initialize.DB.Find(&locacion_usuario).Error; err != nil {
		return locacion_usuario, errors.New("data not found")
	}

	return locacion_usuario, nil
}

func SaveMiniChallengeUsuario(challenges []UsuarioLocacion) error {
	tx := initialize.DB.Begin()

	for _, locacion_usuario := range challenges {

		if err := tx.Create(&UsuarioLocacion{
			Evento:     locacion_usuario.Evento,
			Estado:     locacion_usuario.Estado,
			LocacionId: locacion_usuario.LocacionId,
			UsuarioId:  locacion_usuario.UsuarioId,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}

	}
	return tx.Commit().Error
}

func SaveChallengeUsuario(challenges []UsuarioLocacion) error {
	tx := initialize.DB.Begin()

	for _, locacion_usuario := range challenges {

		if err := tx.Create(&UsuarioLocacion{
			Evento:     locacion_usuario.Evento,
			Estado:     locacion_usuario.Estado,
			LocacionId: locacion_usuario.LocacionId,
			UsuarioId:  locacion_usuario.UsuarioId,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}

	}
	return tx.Commit().Error
}

func RemoveLocaionChallengeUsuario(uid uint) error {
	var err error
	//if err := initialize.DB.Where("UsuarioId = ?", uid).Delete(&UsuarioLocacion{}).Error; err != nil {
	if err := initialize.DB.Where("usuario_id = ? AND  evento = ?", uid, "Challenge").Unscoped().Delete(&UsuarioLocacion{}).Error; err != nil {

		return err
	}
	return err
}

func RemoveAllMiniChallengeUsuario() error {
	var err error
	//if err := initialize.DB.Where("UsuarioId = ?", uid).Delete(&UsuarioLocacion{}).Error; err != nil {
	if err := initialize.DB.Where("evento = ?", "MiniChallenge").Unscoped().Delete(&UsuarioLocacion{}).Error; err != nil {

		return err
	}
	return err
}

func (locacion_usuario *UsuarioLocacion) UpdateLocaionUsuario(uid uint) (*UsuarioLocacion, error) {

	err := initialize.DB.Save(locacion_usuario).Where("ID=?", uid).Error

	if err != nil {
		return &UsuarioLocacion{}, errors.New("user not found")
	}
	return locacion_usuario, err
}

func GetLocaionUsuarioByUsuarioID(uid uint, evento string) (UsuarioLocacion, error) {
	var locacion_usuario = UsuarioLocacion{}

	if err := initialize.DB.Where(&UsuarioLocacion{UsuarioId: uid, Evento: evento}).First(&locacion_usuario).Error; err != nil {

		return locacion_usuario, err
	}
	return locacion_usuario, nil
}
