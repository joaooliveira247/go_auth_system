package schemas

import (
	"encoding/json"

	"github.com/joaooliveira247/go_auth_system/src/errors"
	"github.com/joaooliveira247/go_auth_system/src/security"
)

type UserSchemaIn struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"-"`
}

func (schema *UserSchemaIn) ToCacheJson() (string, error) {
	hashedPassword, err := security.GenHash(schema.Password)

	if err != nil {
		return "", errors.NewCacheError(err)
	}

	schema.Password = hashedPassword

	schemaJson, err := json.Marshal(schema)

	if err != nil {
		return "", errors.NewCacheError(err)
	}

	return string(schemaJson), nil
}
