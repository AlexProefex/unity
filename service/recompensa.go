package service

import (
	"errors"
	"fmt"
	"unity/repository/dao"
	"unity/types"
)

func ServiceGetAllRecompensa() ([]dao.Recompensa, error) {
	recompensa, err := dao.GetAllRecompensas()
	return recompensa, err
}

func ServiceSaveRecompensa(input types.RecompensaRegister) (dao.Recompensa, error) {
	recompensa := dao.Recompensa{
		InsigniaId: input.Insignia,
		UsuarioId:  input.Usuario,
	}
	current, err := recompensa.SaveRecompensa()
	return *current, err
}

func ServiceUpdateRecompensa(input types.RecompensaUpdate, id uint) (dao.Recompensa, error) {

	recompensa := dao.Recompensa{
		ID:         input.ID,
		InsigniaId: input.Insignia,
		UsuarioId:  input.Usuario,
	}
	if recompensa.ID != id {
		return recompensa, errors.New("no se pudo actualizar el recurso solicitado")
	}
	current, err := recompensa.UpdateRecompensa(id)
	return *current, err
}

func ServiceGetRecompensaByID(uid uint) (dao.Recompensa, error) {
	recompensa := dao.Recompensa{
		ID: uid,
	}
	recompensa, err := recompensa.GetRecomenpsaByID(uid)
	return recompensa, err
}

func ServiceGetAllRecompensaByUserId(uid uint) ([]dao.Recompensa, error) {
	recompensa, err := dao.GetAllRecompensasByUserId(uid)
	return recompensa, err
}

func ServiceCobrarAgregarRecompensaInsignia(uid uint, cantidad int, producto uint) (dao.Premio, error) {

	fmt.Println("Lllegamos al servicio")

	response := dao.Premio{}
	recompensa := dao.Usuarios{
		ID:       uid,
		Cantidad: cantidad,
	}

	recompensa, err := recompensa.GetUserByID(uid)

	if err != nil {
		return response, err
	}

	recompensa.Cantidad = recompensa.Cantidad - cantidad

	fmt.Println("Lllegamos al servicio")

	if recompensa.Cantidad < 0 {
		fmt.Println("no cuentas con la cantidad de puntos suficientes")
		return response, errors.New("no cuentas con la cantidad de puntos suficientes")

	}

	if recompensa.ID != uid {
		return response, errors.New("no se pudo actualizar el recurso solicitado")
	}
	_, err = recompensa.CobrarAgregarRecompensaInsignia(uid)
	if err != nil {
		return response, err
	}
	fmt.Println("llego aqui")
	prima := dao.Premio{}
	response, err = prima.GetPremioByID(uint(producto))
	if err != nil {
		return response, err
	}

	return response, err
}

func ServiceCobrarAgregarRecompensaPuntos(uid uint, puntos int, producto uint) (dao.Premio, error) {

	response := dao.Premio{}

	recompensa := dao.Usuarios{
		ID:     uid,
		Puntos: puntos,
	}

	recompensa, err := recompensa.GetUserByID(uid)

	if err != nil {
		return response, err
	}

	recompensa.Puntos = recompensa.Puntos - puntos

	if recompensa.Cantidad < 0 {
		fmt.Println("no cuentas con la cantidad de puntos suficientes")
		return response, errors.New("no cuentas con la cantidad de puntos suficientes")

	}

	if recompensa.ID != uid {
		return response, errors.New("no se pudo actualizar el recurso solicitado")
	}

	_, err = recompensa.CobrarAgregarRecompensaPuntos(uid)
	if err != nil {
		return response, err
	}

	prima := dao.Premio{}
	response, err = prima.GetPremioByID(uint(producto))
	if err != nil {
		return response, err
	}

	return response, err
}

func ServiceGenerarPagoCodigoQR(input types.UsuariosValidateCP, uid uint) (string, error) {

	user := dao.Usuarios{
		ID:       uid,
		Puntos:   input.Puntos,
		Cantidad: input.Cantidad,
	}
	user, err := user.GetUserByID(uid)

	if err != nil {
		return "", err
	}

	user.Cantidad = user.Cantidad - input.Cantidad
	user.Puntos = user.Puntos - input.Puntos

	if user.Cantidad < 0 || user.Puntos < 0 {
		fmt.Println("no cuentas con la cantidad de puntos suficientes")
		return "", errors.New("no cuentas con la cantidad de puntos suficientes")

	}
	fmt.Println(input.Producto)

	token, err := dao.GenerateQRToken(uid, input.Cantidad, input.Puntos, input.Producto)

	return token, err

}
