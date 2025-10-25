package schemas_test

import (
	"fmt"
	"testing"

	"github.com/joaooliveira247/go_auth_system/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestToCacheJsonSuccess(t *testing.T) {

	schema := mocks.GenFakeUserSchemaIn()

	schemaJson, err := schema.ToCacheJson()

	assert.NoError(t, err)
	assert.Contains(t, schemaJson, fmt.Sprintf(`"email":"%s"`, schema.Email))
}
