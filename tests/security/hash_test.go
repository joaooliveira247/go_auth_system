package security_test

import (
	"testing"

	"github.com/joaooliveira247/go_auth_system/src/security"
	"github.com/joaooliveira247/go_auth_system/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGenHashSuccess(t *testing.T) {
	passwd := mocks.GenFakePassword()
	hash, err := security.GenHash(passwd)

	assert.NoError(t, err)
	assert.Len(t, hash, 60)
}

func TestCheckPasswordSuccess(t *testing.T) {
	passwd := mocks.GenFakePassword()

	hash, _ := security.GenHash(passwd)

	err := security.CheckPassword(hash, passwd)

	assert.NoError(t, err)
}
