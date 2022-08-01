package db

import (
	"fmt"
	"layout-item/pkg/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// InitMysqlDB 初始化mysql数据库
func InitMysqlDB(user, pass, addr, db string, slog bool) {
	logger.Logger.Info("init mysql")

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, pass, addr, db)

	var err error
	var newLogger gormlogger.Interface

	if !slog { // false仅控制台显示,true则控制台显示,并记录日志文件
		newLogger = gormlogger.Default.LogMode(gormlogger.Info) // gorm提供的默认日志器
	} else {
		// 使用自定义日志器
		newLogger = gormlogger.New(
			logger.StdLogger(logger.Logger),
			gormlogger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  gormlogger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			},
		)
	}

	DB, err = gorm.Open(mysql.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	logger.Logger.Info("init mysql ok")

	return
}

// 初始模型
type GormModel struct {
	ID        int64          `gorm:"primarykey"`
	CreatedAt int64          `gorm:"autoCreateTime"` // 使用时间戳(秒),避免时区存储问题(time会存储时区) // autoUpdateTime:nano 存储纳秒 :milli 存储毫秒
	UpdatedAt int64          `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Pagination 分页
func Pagination(page, limit uint32) func(db *gorm.DB) *gorm.DB {
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit // 计算偏移量

		return func(db *gorm.DB) *gorm.DB {
			return db.Offset(int(offset)).Limit(int(limit))
		}
	}

	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}
