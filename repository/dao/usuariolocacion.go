package dao

import (
	"errors"
	"unity/initialize"
	"unity/repository/model"
	"unity/utils"
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

func SaveAndDropChallengeUsuario(uid uint, challenges []UsuarioLocacion) error {
	tx := initialize.DB.Begin()
	if err := tx.Where("usuario_id = ? AND  evento = ?", uid, utils.EventChallenge).Unscoped().Delete(&UsuarioLocacion{}).Error; err != nil {
		return errors.New(utils.Ha_ocurrido_un_error)
	}
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
	if err := initialize.DB.Where("usuario_id = ? AND  evento = ?", uid, utils.EventChallenge).Unscoped().Delete(&UsuarioLocacion{}).Error; err != nil {
		return errors.New(utils.Ha_ocurrido_un_error)
	}

	return err
}

func RemoveAllMiniChallengeUsuario() error {
	var err error
	if err := initialize.DB.Where("evento = ?", "MiniChallenge").Unscoped().Delete(&UsuarioLocacion{}).Error; err != nil {
		return errors.New(utils.Ha_ocurrido_un_error)
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
		return locacion_usuario, errors.New(utils.Ha_ocurrido_un_error)
	}
	return locacion_usuario, nil
}

func GetAllLocaionUsuarioByUserId(uid uint) ([]UsuarioLocacion, error) {
	var locacion_usuario []UsuarioLocacion
	if err := initialize.DB.Where(&UsuarioLocacion{UsuarioId: uid}).Find(&locacion_usuario).Error; err != nil {
		return locacion_usuario, errors.New(utils.Not_found)
	}
	return locacion_usuario, nil
}

func GetAllLocaionUsuarioByUserIdAndEstado(uid uint) ([]UsuarioLocacion, error) {
	var locacion_usuario []UsuarioLocacion
	if err := initialize.DB.Where(&UsuarioLocacion{UsuarioId: uid, Estado: "Incompleto"}).Find(&locacion_usuario).Error; err != nil {
		return locacion_usuario, errors.New(utils.Not_found)
	}
	return locacion_usuario, nil
}

func ValidateLocationById(uid uint) ([]UsuarioLocacion, error) {
	var locacion_usuario []UsuarioLocacion
	if err := initialize.DB.Where(&UsuarioLocacion{ID: uid, Estado: "Incompleto"}).First(&locacion_usuario).Error; err != nil {
		return locacion_usuario, errors.New(utils.Not_found)
	}
	return locacion_usuario, nil
}

func ActualizarEstado(uid uint, puntos int, userid uint) error {

	tx := initialize.DB.Begin()

	if err := tx.Model(&UsuarioLocacion{}).Where("ID = ?", uid).Update("estado", "Completo").Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&Usuarios{}).Where("ID = ?", userid).Update("puntos", puntos).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func ValidateAsingRoutesById(uid uint, typeCat string) (int64, error) {
	var count int64
	if err := initialize.DB.Where(&UsuarioLocacion{UsuarioId: uid, Evento: typeCat}).Find(&UsuarioLocacion{}).Count(&count).Error; err != nil {
		return count, errors.New(utils.Not_found)
	}
	return count, nil
}

func ValidateAsingComplete(uid uint, typeCat string) (int64, error) {
	var count int64
	if err := initialize.DB.Where(&UsuarioLocacion{UsuarioId: uid, Evento: typeCat, Estado: "Completo"}).Find(&UsuarioLocacion{}).Count(&count).Error; err != nil {
		return count, errors.New(utils.Not_found)
	}
	return count, nil
}
