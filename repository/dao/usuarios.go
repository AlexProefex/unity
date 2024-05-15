package dao

import (
	"errors"
	"fmt"
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
		return user, errors.New("user not found")
	}
	user.Password = ""
	return user, nil
}

func (user Usuarios) GetUserByName(name string) (Usuarios, error) {

	if err := initialize.DB.Where("nombre = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Usuarios{}, errors.New("user not found")
		}
		return Usuarios{}, err
	}

	user.Password = ""

	return user, nil
}

func GetAllUser() ([]Usuarios, error) {

	var users []Usuarios

	if err := initialize.DB.Omit("Password").Find(&users).Error; err != nil {
		return users, errors.New("data not found")
	}

	return users, nil
}

func (usuarios *Usuarios) SaveUsuarios() (*Usuarios, error) {

	usuarios, err := HashPassowrd(usuarios)
	if err != nil {
		return &Usuarios{}, err
	}

	err = initialize.DB.Create(usuarios).Error

	if err != nil {
		_, err = customError.ValidateUnique(err)

		if err != nil {
			return &Usuarios{}, err
		}
	}

	//user.Apellidos = "marcos "
	//fmt.Println(DB)
	//fmt.Println(usuarios)

	return usuarios, err
}

func HashPassowrd(usuarios *Usuarios) (*Usuarios, error) {
	// Turn password into hash

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuarios.Password), bcrypt.DefaultCost)
	if err != nil {
		return usuarios, err
	}

	usuarios.Password = string(hashedPassword)

	//db.Statement.SetColumn("Password", string(hashedPassword))
	// Remove spaces in username
	//usuarios.Nombre = html.EscapeString(strings.TrimSpace(usuarios.Nombre))
	//fmt.Println("CAll HOKKKKKK")

	return usuarios, nil
}

/*
func (usuarios *Usuarios) BeforeCreate(db *gorm.DB) error {
	// Turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuarios.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usuarios.Password = string(hashedPassword)

	// Remove spaces in username
	//	usuarios.Nombre = html.EscapeString(strings.TrimSpace(usuarios.Nombre))

	return nil
}*/

/*
func (user *Usuarios) UpdateUser() (*Usuarios, error) {
	if user.ID == 0 {
		return user, errors.New("User not found!")
	}

	err := initialize.DB.Model(user).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
*/
// VerifyPassword compares the provided password with the hashed password
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
		fmt.Println("passwrod verify failed", err)

		return "", err
	}

	token, err := utils.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

// GenerateQRToken crea un token para realizar los cobros de premios
func GenerateQRToken(uid uint, cantidad int, puntos int) (string, error) {
	var err error

	u := Usuarios{}

	err = initialize.DB.Model(Usuarios{}).Where("ID = ?", uid).Take(&u).Error

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateQrToken(u.ID, cantidad, puntos)

	if err != nil {
		return "", err
	}
	return token, nil

}

func (usuario *Usuarios) CobrarAgregarRecompensaInsignia(uid uint) (*Usuarios, error) {

	err := initialize.DB.Debug().Model(&Usuarios{}).Where("ID = ?", uid).Update("cantidad", usuario.Cantidad).Error

	if err != nil {
		return &Usuarios{}, errors.New("user not found")
	}
	return usuario, err
}

func (usuario *Usuarios) CobrarAgregarRecompensaPuntos(uid uint) (*Usuarios, error) {

	err := initialize.DB.Debug().Model(&Usuarios{}).Where("ID = ?", uid).Update("puntos", usuario.Puntos).Error

	if err != nil {
		return &Usuarios{}, err
	}
	return usuario, err
}

func (usuario *Usuarios) CambiarPassword(uid uint) (*Usuarios, error) {

	usuarios, err := HashPassowrd(usuario)
	if err != nil {
		return &Usuarios{}, errors.New("la contraseña no cumple con los requisitos minimos")
	}

	err = initialize.DB.Model(&Usuarios{}).Where("ID = ?", uid).Update("password", usuarios.Password).Error

	if err != nil {
		return &Usuarios{}, errors.New("no se pudo actualizar la contraseña")
	}
	return &Usuarios{}, err
}

func (user Usuarios) ConsultarEmail(correo_electronico string) (Usuarios, error) {
	if err := initialize.DB.Where("Correo_electronico = ?", correo_electronico).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Usuarios{}, errors.New("user not found")
		}
		return Usuarios{}, err
	}
	return user, nil
}

func (user Usuarios) ConsultarSecret(correo_electronico string, secret string) (Usuarios, error) {

	if err := initialize.DB.Where("Correo_electronico = ? and secret=?", correo_electronico, secret).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Usuarios{}, errors.New("invalid secret key")
		}
		return Usuarios{}, err
	}

	user.Password = ""

	return user, nil
}
