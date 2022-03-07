package db

import (
	"fmt"
	"layout-item/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// InitMysqlDB 初始化mysql数据库
func InitMysqlDB(user, pass, addr, db string) {
	logger.Logger.Info("init mysql")

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, pass, addr, db)

	var err error
	DB, err = gorm.Open(mysql.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		panic(err)
	}

	logger.Logger.Info("init mysql ok")

	return
}
