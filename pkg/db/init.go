package db

import (
	"layout-item/configs"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	RedisCli *redis.Client
)

func Init(cfg *configs.BaseConfig) {
	InitMysqlDB(cfg.MYSQL.User, cfg.MYSQL.Pass, cfg.MYSQL.Addr, cfg.MYSQL.Db, false)
	InitRedis(cfg.Redis.Addr, cfg.Redis.Addr, cfg.Redis.Db)

	return
}
