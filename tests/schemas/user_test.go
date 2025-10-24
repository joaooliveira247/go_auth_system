package schemas_test

import (
	"fmt"
	"testing"

	"github.com/joaooliveira247/go_auth_system/src/schemas"
	"github.com/joaooliveira247/go_auth_system/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestToCacheJsonSuccess(t *testing.T) {

	user := mocks.GenFakeUser()

	schema := schemas.UserSchemaIn{
		Email:           user.Email,
		Password:        user.Password,
		ConfirmPassword: user.Password,
	}

	schemaJson, err := schema.ToCacheJson()

	assert.NoError(t, err)
	assert.Contains(t, schemaJson, fmt.Sprintf(`"email":"%s"`, user.Email))
}
