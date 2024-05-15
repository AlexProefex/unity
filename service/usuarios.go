package service

import (
	"errors"
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

func ServiceRecuperarContrasena(input types.UsuariosPassword) (dao.Usuarios, error) {

	usuarios := dao.Usuarios{
		Correo_electronico: input.Correo_electronico,
		Password:           input.Password,
		Secret:             input.Secret,
	}

	_, err := usuarios.ConsultarEmail(input.Correo_electronico)
	if err != nil {
		return usuarios, err
	}

	usuario, err := usuarios.ConsultarSecret(input.Correo_electronico, input.Secret)
	if err != nil {
		return usuarios, err
	}

	usuario.Password = input.Password
	_, err = usuario.CambiarPassword(usuario.ID)
	if err != nil {
		return usuarios, err
	}

	return usuario, err

}

func ServiceCambiarContrasena(input types.ConfirmPassword) (dao.Usuarios, error) {

	usuarios := dao.Usuarios{
		Correo_electronico: input.Correo_electronico,
		Password:           input.Password,
	}

	usuario, err := usuarios.ConsultarEmail(input.Correo_electronico)
	if err != nil {
		return usuarios, errors.New("el correo no existe")
	}

	err = dao.VerifyPassword(input.Password, usuario.Password)
	if err != nil {
		return usuarios, errors.New("la contraseña ingresada es incorrecta")
	}

	usuario.Password = input.NewPassword

	_, err = usuario.CambiarPassword(usuario.ID)

	return usuarios, err

}
