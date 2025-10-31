package cache

import "github.com/redis/go-redis/v9"

type Cache interface {
	Set(key string, value interface{}) error
}

type redisCache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) Cache {
	return &redisCache{client}
}
