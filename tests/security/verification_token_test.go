package security_test

import (
	"testing"

	"github.com/joaooliveira247/go_auth_system/src/security"
	"github.com/stretchr/testify/assert"
)

func TestGenerateSignUpTokenSuccess(t *testing.T) {
	token, err := security.GenerateSignUpToken(16)

	assert.NoError(t, err)
	assert.Len(t, token, 32)
}
