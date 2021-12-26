package db

import (
	"layout-item/pkg/logger"

	"github.com/go-redis/redis"
)

// InitRedis 初始化Redis
func InitRedis(addr, password string, redisDB int) {
	logger.Logger.Info("init redis")
	redisCli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       redisDB,
	})

	_, err := redisCli.Ping().Result()
	if err != nil {
		panic(err)
	}

	logger.Logger.Info("init redis ok")

	RedisCli = redisCli

	return
}
