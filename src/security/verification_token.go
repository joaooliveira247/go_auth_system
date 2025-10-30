package security

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/joaooliveira247/go_auth_system/src/errors"
)

func GenerateSignUpToken(n int) (string, error) {
	tokenBytes := make([]byte, n)

	if _, err := rand.Read(tokenBytes); err != nil {
		return "", errors.NewTokenSignUpError(err)
	}

	return hex.EncodeToString(tokenBytes), nil
}
