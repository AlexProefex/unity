package service

import (
	"errors"
	"fmt"
	"unity/repository/dao"
	"unity/types"
	"unity/utils"
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
		Secret:             input.Secret,
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
		return usuarios, errors.New("not found")
	}
	usuario, err := usuarios.ConsultarSecret(input.Correo_electronico, input.Secret)
	if err != nil {
		return usuarios, errors.New("failed secret key")
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
		return usuarios, errors.New("not found")
	}
	err = dao.VerifyPassword(input.Password, usuario.Password)
	if err != nil {
		return usuarios, errors.New(utils.VerifyPassword)
	}
	usuario.Password = input.NewPassword
	_, err = usuario.CambiarPassword(usuario.ID)
	return usuarios, err
}

func ServiceUpdatePerfil(input types.UpdatePerfil, uid uint) (dao.Usuarios, error) {
	usuarios := dao.Usuarios{
		Nombre:    input.Nombre,
		Apellidos: input.Apellidos,
	}
	usuario, err := usuarios.UpdatePerfil(uid)
	if err != nil {
		return *usuario, err
	}
	return *usuario, err
}

func ServiceReclamarPuntos(input types.AsingarPuntos, uid uint) (string, error) {
	count, err := dao.ValidateAsingRoutesById(uid, input.TypeRoute)
	if err != nil {
		return "dao.Usuarios{}", err
	}
	if count < 1 {
		return "dao.Usuarios{}", err
	}

	usuario := dao.Usuarios{}
	_, err = dao.ValidateLocationById(input.ID)
	if err != nil {
		return "dao.Usuarios{}", err
	}

	user, err := usuario.GetUserByID(uid)

	if err != nil {
		return "dao.Usuarios{}", err
	}

	user.Puntos = user.Puntos + input.Puntos

	err = dao.ActualizarEstado(input.ID, user.Puntos, uid)

	if err != nil {
		return "dao.Usuarios{}", err
	}

	cantidad, err := dao.ValidateAsingComplete(uid, input.TypeRoute)
	fmt.Println(user.Cantidad, cantidad, count)
	if cantidad == count {
		user.Cantidad = user.Cantidad + 1
		dao.UpdateInsignia(uid, user.Cantidad, input.Insignia)
		return "completo", err
	}

	return "dao.Usuarios{}", err
}
