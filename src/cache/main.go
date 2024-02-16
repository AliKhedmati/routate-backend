package cache

import (
	"github.com/AliKhedmati/routate-backend/config"
	cacheDriver "github.com/AliKhedmati/routate-backend/src/cache/drivers"
	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Connect() error
	Close() error
	GetCache() *redis.Client
}

var (
	configs *config.Config
	cache   Cache
)

func Init() error {
	var err error
	configs = config.GetInstance()
	cache = &cacheDriver.Redis{
		Host:     configs.Get("REDIS_HOST"),
		Port:     configs.Get("REDIS_PORT"),
		Password: configs.Get("REDIS_PASSWORD"),
	}

	// Connect to the database
	if err = cache.Connect(); err != nil {
		return err
	}

	return err
}

func GetCache() *redis.Client {
	return cache.GetCache()
}
