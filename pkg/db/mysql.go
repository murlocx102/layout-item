package db

import (
	"fmt"
	"layout-item/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// InitMysqlDB 初始化mysql数据库
func InitMysqlDB(user, pass, addr, db string) {
	logger.Logger.Info("init mysql")

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, pass, addr, db)

	var err error
	DB, err = gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	DB.SingularTable(true)
	DB.LogMode(true)

	logger.Logger.Info("init mysql ok")

	return
}
