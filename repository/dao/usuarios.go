package dao

import (
	"errors"
	"fmt"
	"strconv"
	"unity/customError"
	"unity/initialize"
	"unity/repository/model"
	"unity/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Usuarios model.Usuarios

func (user Usuarios) GetUserByID(uid uint) (Usuarios, error) {

	if err := initialize.DB.First(&user, uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New(utils.Not_found)
		}
		return user, errors.New(utils.Ha_ocurrido_un_error)
	}
	user.Password = ""
	return user, nil
}

/*
func (user Usuarios) GetUserByName(name string) (Usuarios, error) {

	if err := initialize.DB.Where("nombre = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Usuarios{}, errors.New("user not found")
		}
		return Usuarios{}, err
	}

	user.Password = ""

	return user, nil
}*/

func GetAllUser() ([]Usuarios, error) {
	var users []Usuarios
	if err := initialize.DB.Omit("Password").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, errors.New(utils.Not_found)
		}
	}
	return users, nil
}

func (usuarios *Usuarios) SaveUsuarios() (*Usuarios, error) {

	usuarios, err := HashPassowrd(usuarios)
	if err != nil {
		return &Usuarios{}, errors.New(utils.PasswordError)
	}
	err = initialize.DB.Create(usuarios).Error
	if err != nil {
		_, err = customError.ValidateUnique(err)
		if err != nil {
			return &Usuarios{}, err
		}
	}
	return usuarios, err
}

func HashPassowrd(usuarios *Usuarios) (*Usuarios, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuarios.Password), bcrypt.DefaultCost)
	if err != nil {
		return usuarios, errors.New(utils.PasswordError)
	}
	usuarios.Password = string(hashedPassword)
	return usuarios, nil
}

func VerifyPassword(password, hashedPassword string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// LoginCheck validates user credentials and generates a token
func LoginCheck(correo_electrónico, password string) (string, error) {
	var err error
	u := Usuarios{}
	err = initialize.DB.Model(Usuarios{}).Where("correo_electronico = ?", correo_electrónico).Take(&u).Error
	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", errors.New(utils.VerifyPassword)
	}
	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		return "", errors.New(utils.TokenError)
	}
	return token, nil
}

// GenerateQRToken crea un token para realizar los cobros de premios
func GenerateQRToken(uid uint, cantidad int, puntos int, producto int) (string, error) {
	var err error
	u := Usuarios{}
	err = initialize.DB.Model(Usuarios{}).Where("ID = ?", uid).Take(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New(utils.Not_found)
		}
		return "", err
	}
	fmt.Println("QR SERvice ", producto)
	token, err := utils.GenerateQrToken(u.ID, cantidad, puntos, producto)
	if err != nil {
		return "", errors.New(utils.Ha_ocurrido_un_error)
	}
	return token, nil
}

func (usuario *Usuarios) CobrarAgregarRecompensaInsignia(uid uint) (*Usuarios, error) {
	err := initialize.DB.Model(&Usuarios{}).Where("ID = ?", uid).Update("cantidad", usuario.Cantidad).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Usuarios{}, errors.New(utils.Not_found)
		}
		return &Usuarios{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	return usuario, err
}

func (usuario *Usuarios) CobrarAgregarRecompensaPuntos(uid uint) (*Usuarios, error) {
	err := initialize.DB.Model(&Usuarios{}).Where("ID = ?", uid).Update("puntos", usuario.Puntos).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Usuarios{}, errors.New(utils.Not_found)
		}
		return &Usuarios{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	return usuario, err
}

func CobrarAgregarRecompensaInsigniaWithNewCamps(uid uint, cantidad int, producto int, insignias int) (*Usuarios, error) {
	tx := initialize.DB.Begin()

	usuario := &Usuarios{}

	err := initialize.DB.Model(usuario).Where("ID = ?", uid).Update("cantidad", cantidad).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return &Usuarios{}, errors.New(utils.Not_found)
		}
		tx.Rollback()
		return &Usuarios{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	canjes := &Canjes{
		ProductId: strconv.Itoa(producto),
		UsuarioId: strconv.Itoa(int(uid)),
		Puntos:    "0",
		Insignias: strconv.Itoa(insignias),
	}

	err = initialize.DB.Create(canjes).Error
	if err != nil {
		_, err = customError.ValidateUnique(err)
		if err != nil {
			return &Usuarios{}, err
		}
	}

	tx.Commit()

	return usuario, err

}

func CobrarAgregarRecompensaPuntosWithNewCamps(uid uint, puntos int, producto int, points int) (*Usuarios, error) {
	tx := initialize.DB.Begin()

	usuario := &Usuarios{}

	err := initialize.DB.Model(usuario).Where("ID = ?", uid).Update("puntos", puntos).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return &Usuarios{}, errors.New(utils.Not_found)
		}
		tx.Rollback()
		return &Usuarios{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	canjes := &Canjes{
		ProductId: strconv.Itoa(producto),
		UsuarioId: strconv.Itoa(int(uid)),
		Puntos:    strconv.Itoa(points),
		Insignias: "0",
	}

	err = initialize.DB.Create(canjes).Error
	if err != nil {
		_, err = customError.ValidateUnique(err)
		if err != nil {
			return &Usuarios{}, err
		}
	}

	tx.Commit()

	return usuario, err

}

func (usuario *Usuarios) CambiarPassword(uid uint) (*Usuarios, error) {
	usuarios, err := HashPassowrd(usuario)
	if err != nil {
		return &Usuarios{}, errors.New(utils.PasswordError)
	}
	err = initialize.DB.Model(&Usuarios{}).Where("ID = ?", uid).Update("password", usuarios.Password).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Usuarios{}, errors.New(utils.Not_found)
		}
		return &Usuarios{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	return &Usuarios{}, err
}

func (user Usuarios) ConsultarEmail(correo_electronico string) (Usuarios, error) {
	if err := initialize.DB.Where("Correo_electronico = ?", correo_electronico).First(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Usuarios{}, errors.New(utils.Not_found)
		}
		return Usuarios{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	return user, nil
}

func (user Usuarios) ConsultarSecret(correo_electronico string, secret string) (Usuarios, error) {
	if err := initialize.DB.Where("Correo_electronico = ? and secret=?", correo_electronico, secret).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Usuarios{}, errors.New(utils.Not_found)
		}
		return Usuarios{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	user.Password = ""
	return user, nil
}

func (usuario *Usuarios) UpdatePerfil(uid uint) (*Usuarios, error) {
	err := initialize.DB.Model(&Usuarios{}).Where("ID = ?", uid).Updates(Usuarios{Nombre: usuario.Nombre, Apellidos: usuario.Apellidos}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &Usuarios{}, errors.New(utils.Not_found)
		}
		return &Usuarios{}, errors.New(utils.Ha_ocurrido_un_error)
	}
	return usuario, err
}

func UpdateInsignia(uid uint, cantidad int, insignia uint) error {
	tx := initialize.DB.Begin()
	err := tx.Model(&Usuarios{}).Where("ID = ?", uid).Update("Cantidad", cantidad).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Save(&Recompensa{UsuarioId: uid, InsigniaId: insignia}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
