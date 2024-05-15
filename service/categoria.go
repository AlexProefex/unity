package service

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
	"unity/repository/dao"
	"unity/types"
	"unity/utils"

	"gorm.io/gorm"
)

func ServiceGetAllCategoria() ([]dao.Categoria, error) {
	categoria, err := dao.GetAllCategorias()
	return categoria, err
}

func ServiceSaveCategoria(input types.CategoriaRegister) (dao.Categoria, error) {
	categoria := dao.Categoria{
		Nombre:      input.Nombre,
		Descripcion: input.Descripcion,
		Insignia:    input.Insignia,
		Tiempo:      input.Tiempo,
	}
	current, err := categoria.SaveCategorias()
	return *current, err
}

func ServiceUpdateCategoria(input types.CategoriaUpdate, id uint) (dao.Categoria, error) {

	categoria := dao.Categoria{
		ID:          input.ID,
		Nombre:      input.Nombre,
		Descripcion: input.Descripcion,
		Insignia:    input.Insignia,
		Tiempo:      input.Tiempo,
	}

	if categoria.ID != id {
		return categoria, errors.New("no se pudo actualizar el recurso solicitado dgfgdgd")
	}

	current, err := categoria.UpdateCategorias(id)
	return *current, err
}

func ServiceGetCAtegoriaByID(uid uint) (dao.Categoria, error) {
	categoria := dao.Categoria{
		ID: uid,
	}
	categoria, err := categoria.GetCategoriaByID(uid)
	return categoria, err
}

func ServiceGetChallenge(input types.CategoriaChallenge) ([]dao.Categoria, error) {
	categoria, err := dao.GetChallenge(input.ID)
	return categoria, err
}

func ServiceSetChallenge(input types.CategoriaChallenge) error {

	var err error
	categoria, err := dao.GetChallenge(input.ID)
	if err != nil {
		return err
	}

	locacion_usuario, err := dao.GetLocaionUsuarioByUsuarioID(input.Usuario, "Challenge")
	isEmpty := false
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			isEmpty = true
		} else {
			return err
		}
	}

	if isEmpty {
		err = createLocationOnUser(input, categoria, "Challenge")
		if err != nil {
			return err
		}

	} else {

		challengeRenue := utils.GetHours(input.Tiempo, locacion_usuario.CreatedAt.UTC())

		if challengeRenue {

			err = dao.RemoveLocaionChallengeUsuario(input.Usuario)

			if err != nil {
				return err
			}
			err = createLocationOnUser(input, categoria, "Challenge")
			if err != nil {
				return err
			}
		}

	}
	return err

}

func removeIndex(s []dao.Locacion, index int) []dao.Locacion {
	return append(s[:index], s[index+1:]...)
}

func ServiceSetMiniChallenge(input types.CategoriaChallenge) error {

	var err error
	locacion, err := dao.GetAllLocacion()
	if err != nil {
		return err
	}
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for len(locacion) >= 6 {
		index := rng.Intn(len(locacion)-1) + 1
		locacion = removeIndex(locacion, index)
		fmt.Println(locacion)
	}

	_, err = dao.GetLocaionUsuarioByUsuarioID(input.Usuario, "MiniChallenge")

	fmt.Println("llama al meotdo")
	isEmpty := false
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			isEmpty = true
		} else {
			return err
		}
	}

	if isEmpty {
		err = createMiniLocationOnUser(input, locacion, "MiniChallenge")
		if err != nil {
			return err
		}

	}
	return err

}

func ServiceRemoveMiniChallenge() error {
	err := dao.RemoveAllMiniChallengeUsuario()
	return err

}

func createLocationOnUser(input types.CategoriaChallenge, categoria []dao.Categoria, evento string) error {
	var err error
	var locaciones []dao.UsuarioLocacion
	for _, row := range categoria[0].Locaciones {

		locaciones = append(locaciones, dao.UsuarioLocacion{
			UsuarioId:  input.Usuario,
			LocacionId: row.ID,
			Evento:     evento,
			Estado:     "Incompleto",
		})
	}
	err = dao.SaveChallengeUsuario(locaciones)
	if err != nil {
		return err
	}
	return err

}

func createMiniLocationOnUser(input types.CategoriaChallenge, locacion []dao.Locacion, evento string) error {
	var err error
	var locaciones []dao.UsuarioLocacion
	for _, row := range locacion {

		locaciones = append(locaciones, dao.UsuarioLocacion{
			UsuarioId:  input.Usuario,
			LocacionId: row.ID,
			Evento:     evento,
			Estado:     "Incompleto",
		})
	}
	err = dao.SaveChallengeUsuario(locaciones)
	if err != nil {
		return err
	}
	return err

}
