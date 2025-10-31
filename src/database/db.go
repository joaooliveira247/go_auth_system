package database

import (
	"github.com/joaooliveira247/go_auth_system/src/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(config.DBURL),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetCacheConnection() (*redis.Client, error) {

	opt, err := redis.ParseURL(config.CacheUrl)

	if err != nil {
		return nil, err
	}

	return redis.NewClient(opt), nil

}
