package customError

import (
	"errors"
	"strings"
	"unity/utils"
)

type UniqueError struct{}

func (m *UniqueError) Error() string {
	return utils.Duplicate_key
}

func ValidateUnique(err error) (string, error) {
	if strings.Contains(err.Error(), "Error 1062 (23000): Duplicate entry") {
		return "", &UniqueError{}
	}
	return "", errors.New(utils.Ha_ocurrido_un_error)
}
