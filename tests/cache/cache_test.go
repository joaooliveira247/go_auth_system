package cache

import (
	"fmt"
	"testing"
	"time"

	"github.com/joaooliveira247/go_auth_system/src/cache"
	"github.com/joaooliveira247/go_auth_system/src/security"
	"github.com/joaooliveira247/go_auth_system/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCacheSetSuccess(t *testing.T) {
	cacheDB, mock := mocks.SetupMockCache()

	user := mocks.GenFakeUserSchemaIn()

	userJson, _ := user.ToCacheJson()

	userSignUpToken, _ := security.GenerateSignUpToken(16)

	mock.ExpectSet(fmt.Sprintf("signup:%s", userSignUpToken), userJson, time.Second*7200).SetVal("OK")

	cache := cache.NewCache(cacheDB)

	err := cache.Set(fmt.Sprintf("signup:%s", userSignUpToken), userJson)

	assert.NoError(t, err)

}
