package service

import (
	"unity/repository/dao"
	"unity/types"
)

func ServiceRegister(input types.UsuariosRegister) (dao.Usuarios, error) {
	user := dao.Usuarios{
		Nombre:             input.Nombre,
		Apellidos:          input.Apellidos,
		Correo_electronico: input.Correo_electronico,
		Codigo_pais:        input.Codigo_pais,
		Celular:            input.Celular,
		Genero:             input.Genero,
		Fecha_nacimiento:   input.Fecha_nacimiento,
		Nacionalidad:       input.Nacionalidad,
		Password:           input.Password,
	}
	current, err := user.SaveUsuarios()
	return *current, err
}

func ServiceLogin(input types.UsuariosLogin) (string, error) {
	token, err := dao.LoginCheck(input.Correo_electronico, input.Password)
	return token, err
}

func ServiceGetUserByID(id uint) (dao.Usuarios, error) {
	user := dao.Usuarios{
		ID: id,
	}
	user, err := user.GetUserByID(id)
	return user, err
}

func ServiceGetAllUser() ([]dao.Usuarios, error) {

	user, err := dao.GetAllUser()
	return user, err
}
