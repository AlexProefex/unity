package customError

import "strings"

type UniqueError struct{}

func (m *UniqueError) Error() string {
	return "El correo ya se encuentra registrado"
}

func ValidateUnique(err error) (string, error) {
	if strings.Contains(err.Error(), "Error 1062 (23000): Duplicate entry") {
		return "", &UniqueError{}
	}
	return "", err
}
