package main

import (
	merchant "layout-item/api/merchant_demo"
	"layout-item/api/merchant_demo/vip/handler"
	"layout-item/configs"
	"layout-item/pkg/logger"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	Version string
)

func main() {
	log.Println("version:", Version)
	err := configs.LoadConfig("", "")
	if err != nil {
		panic("加载配置文件失败" + err.Error())
	}
	cfg := configs.GetConfig()

	loggerConf := logger.Conf{}.DefaultConf()
	logger.NewLogger(&loggerConf, "merchant") // 初始化日志器

	InitServer(logger.Logger, cfg.MerchantHTTP)
}

// InitServer 初始化服务
func InitServer(logger *zap.Logger, cfg configs.HTTP) {
	e := gin.Default()

	handler.NewVipHttpHandler()

	merchant.Init(e) // 路由
	e.Run(":" + strconv.Itoa(int(cfg.Port)))
}
