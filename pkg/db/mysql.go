package db

import (
	"fmt"
	"layout-item/pkg/logger"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type GormModel struct {
	ID        int64          `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// InitMysqlDB 初始化mysql数据库
func InitMysqlDB(user, pass, addr, db string, slog bool) *gorm.DB {
	logger.Logger.Info("init mysql")

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", user, pass, addr, db) // 使用本地时区(loc=Local)

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

	conn, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	logger.Logger.Info("init mysql ok")

	return conn
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

// Order slice可多条件 []map[SortField]SortMode  SortMode: "desc","2"表示倒序
func Order(sortElems []map[string]string) func(db *gorm.DB) *gorm.DB {
	if len(sortElems) == 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}

	return func(db *gorm.DB) *gorm.DB {
		for _, elems := range sortElems {
			for key, value := range elems {
				switch value {
				case "desc":
					db = db.Order(strings.Join([]string{key, "DESC"}, " "))
				case "2": // "1":正序 "2":倒序
					db = db.Order(strings.Join([]string{key, "DESC"}, " "))
				default:
					db = db.Order(strings.Join([]string{key, "ASC"}, " "))
				}
			}
		}
		return db
	}
}
