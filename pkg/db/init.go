package db

import (
	"context"
	"errors"
	"fmt"
	"layout-item/configs"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
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
	//	NewDBConnMap(cfg.MYSQL)

	return
}

// 以下用于适配商户系统
type DbConnI interface {
	GinWithDBConn(ctx *gin.Context) *gorm.DB

	SetCtxWithDB(ctx context.Context, merchantID int) context.Context
	GetCtxDBConn(ctx context.Context) *gorm.DB
}

type DbConnMap struct {
	defaultCfg configs.MYSQL
	connMap    sync.Map
	ctxDBKey   string
	ctxIDKey   string
}

func NewDBConnMap(cfg configs.MYSQL) *DbConnMap {
	return &DbConnMap{
		defaultCfg: cfg,
		connMap:    sync.Map{},
		ctxDBKey:   "dbConn",
		ctxIDKey:   "merchantID",
	}
}

// 构建库名
func (d *DbConnMap) NewDBName(merchantID int) string {
	if merchantID <= 0 {
		return d.defaultCfg.Db
	}

	return strings.Join([]string{d.defaultCfg.Db, strconv.Itoa(merchantID)}, "_")
}

// 适用于gin.context 响应:商户dbConn句柄
func (d *DbConnMap) GinWithDBConn(ctx *gin.Context) *gorm.DB {
	merchantID, err := d.ginWithMerID(ctx)
	if merchantID == 0 || err != nil {
		fmt.Println("未获获取到商户ID", err.Error())
		return nil
	}

	dbCtx := d.SetCtxWithDB(ctx, int(merchantID))
	return d.GetCtxDBConn(dbCtx)
}

// 适用于gin.context 响应:商户ID
func (d *DbConnMap) ginWithMerID(ctx *gin.Context) (int64, error) {
	value, _ := ctx.Get("merchantID") // 此key对应auth中的token.subjectKey
	merchantID, err := strconv.ParseInt(value.(string), 10, strconv.IntSize)
	if merchantID == 0 || err != nil {
		fmt.Println("未获获取到商户ID", err.Error())
		return 0, errors.New("未获获取到商户ID")
	}

	return merchantID, nil
}

// 原始context.Context设置db信息
func (d *DbConnMap) SetCtxWithDB(ctx context.Context, merchantID int) context.Context {
	dbName := d.NewDBName(merchantID)

	var tx *gorm.DB
	if conn, ok := d.connMap.Load(dbName); ok {
		tx = conn.(*gorm.DB)
	} else {
		tx = InitMysqlDB(d.defaultCfg.User, d.defaultCfg.Pass, d.defaultCfg.Addr, d.defaultCfg.Db, d.defaultCfg.Slog)
		d.connMap.Store(dbName, tx)
	}

	if ctx == nil {
		ctx = context.Background()
	}
	dbCtx := context.WithValue(ctx, d.ctxDBKey, tx)
	return context.WithValue(dbCtx, d.ctxIDKey, merchantID)
}

// 获取ctx中db的连接句柄
func (d *DbConnMap) GetCtxDBConn(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(d.ctxDBKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}

	return nil // 不返回基础db连接句柄,预防错误导致变更基础数据
}
