package models_test

import (
	"testing"

	"github.com/joaooliveira247/go_auth_system/src/models"
	"github.com/joaooliveira247/go_auth_system/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewUserModelFromCacheSuccess(t *testing.T) {
	mockSchema := mocks.GenFakeUserSchemaIn()

	mockCache, _ := mockSchema.ToCacheJson()

	user, err := models.NewUserModelFromCache(mockCache)

	assert.NoError(t, err)

	assert.Equal(t, mockSchema.Email, user.Email)
}
