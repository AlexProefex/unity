package service

import (
	"errors"
	"math/rand"
	"time"
	"unity/repository/dao"
	"unity/types"
	"unity/utils"
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
		return categoria, errors.New(utils.InvalidID)
	}
	current, err := categoria.UpdateCategorias(id)
	return *current, err
}

func ServiceGetCategoriaByID(uid uint) (dao.Categoria, error) {
	categoria := dao.Categoria{}
	categoria, err := categoria.GetCategoriaByID(uid)
	return categoria, err
}

func ServiceGetChallenge(uid uint) ([]dao.Categoria, error) {
	categoria, err := dao.GetChallenge(uid)
	return categoria, err
}

func ServiceSetChallenge(input types.CategoriaChallenge) error {
	categoria, err := dao.GetChallenge(input.ID)
	if err != nil {
		return err
	}
	count, err := dao.ValidateAsingRoutesById(input.Usuario, utils.EventChallenge)
	if err != nil {
		return err
	}
	if count < 1 {
		locaciones := createLocationOnUser(input, categoria, utils.EventChallenge)
		err = dao.SaveChallengeUsuario(locaciones)
		if err != nil {
			return err
		}
	} else {
		locacion_usuario, err := dao.GetLocaionUsuarioByUsuarioID(input.Usuario, utils.EventChallenge)
		if err != nil {
			return err
		}
		challengeRenue := utils.GetHours(input.Tiempo, locacion_usuario.CreatedAt.UTC())
		if challengeRenue {
			locaciones := createLocationOnUser(input, categoria, utils.EventChallenge)
			err = dao.SaveAndDropChallengeUsuario(input.Usuario, locaciones)
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
	locacion, err := dao.GetAllLocacion()
	if err != nil {
		return err
	}
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for len(locacion) >= 6 {
		index := rng.Intn(len(locacion)-1) + 1
		locacion = removeIndex(locacion, index)
	}
	count, err := dao.ValidateAsingRoutesById(input.Usuario, utils.EventMiniChallenge)
	if err != nil {
		return err
	}
	if count < 1 {
		err = createMiniLocationOnUser(input, locacion, utils.EventMiniChallenge)
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

func createLocationOnUser(input types.CategoriaChallenge, categoria []dao.Categoria, evento string) []dao.UsuarioLocacion {
	var locaciones []dao.UsuarioLocacion
	for _, row := range categoria[0].Locaciones {
		locaciones = append(locaciones, dao.UsuarioLocacion{
			UsuarioId:  input.Usuario,
			LocacionId: row.ID,
			Evento:     evento,
			Estado:     utils.StatusIncomplete,
		})
	}

	return locaciones
}

func createMiniLocationOnUser(input types.CategoriaChallenge, locacion []dao.Locacion, evento string) error {
	var err error
	var locaciones []dao.UsuarioLocacion
	for _, row := range locacion {
		locaciones = append(locaciones, dao.UsuarioLocacion{
			UsuarioId:  input.Usuario,
			LocacionId: row.ID,
			Evento:     evento,
			Estado:     utils.StatusIncomplete,
		})
	}
	err = dao.SaveChallengeUsuario(locaciones)
	if err != nil {
		return err
	}
	return err
}
