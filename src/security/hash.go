package security

import (
	"github.com/joaooliveira247/go_auth_system/src/errors"
	"golang.org/x/crypto/bcrypt"
)

func GenHash(passwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.NewHashError(err)
	}

	return string(hash), nil
}
