package main

import (
	"fmt"
	"layout-item/configs"
	"layout-item/infrastructure/logger"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

var (
	Version string
)

func main() {
	log.Println("version:", Version)
	err := configs.LoadConfing("", "")
	if err != nil {
		panic("加载配置文件失败" + err.Error())
	}
	cfg := configs.GetConfig()

	loggerConf := logger.Conf{}.DefaultConf()
	logger.NewLogger(&loggerConf, "live") // 初始化日志器

	logger.Logger.Info("读取配置参数测试", zap.Uint("端口", cfg.HTTP.Port))

	//db.Init(cfg) // 初始化db

	fmt.Println("hello,world!")

	shutDown()
}

// shutDown 处理应用关闭
func shutDown() {
	c := make(chan os.Signal, 1)
	// windows 不支持  syscall.SIGTSTP, syscall.SIGSTOP 信号，需要考虑实现
	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	<-c

	fmt.Println("收到关闭信号")
}

//应用选项
type Option struct {
}
