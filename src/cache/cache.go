package cache

import (
	"context"
	"time"

	"github.com/joaooliveira247/go_auth_system/src/config"
	"github.com/joaooliveira247/go_auth_system/src/errors"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(key string, value interface{}) error
}

type redisCache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) Cache {
	return &redisCache{client}
}

func (redisCache *redisCache) Set(key string, value interface{}) error {
	ctx := context.Background()
	if err := redisCache.client.Set(
		ctx,
		key,
		value,
		time.Duration(config.CacheDuration)*time.Second).Err(); err != nil {
		return errors.NewCacheError(err)
	}
	return nil
}
