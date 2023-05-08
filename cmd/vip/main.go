package main

import (
	"fmt"
	"layout-item/configs"
	"layout-item/internal/vip/service"
	"layout-item/internal/vip/service/proto"
	etcdkit "layout-item/pkg/etcd"
	"layout-item/pkg/logger"

	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	ServerName    = "vip"
	ServerAddr    = "127.0.0.1:8083"
	ServerVersion = "1.0.0"
)

func main() {
	err := configs.LoadConfig("../../../configs/", "")
	if err != nil {
		panic("加载配置文件失败" + err.Error())
	}
	cfg := configs.GetConfig()
	server := service.InitUserServer(cfg.MYSQL)

	loggerConf := logger.Conf{}.DefaultConf()
	logger.NewLogger(&loggerConf, ServerName) // 初始化日志器

	go etcdServer(logger.Logger, server)
	//...

	shutDown()
}

func etcdServer(logger *zap.Logger, server *service.VipServer) {
	listen, err := net.Listen("tcp", ServerAddr)
	if err != nil {
		logger.Error("开启端口监听失败", zap.Error(err))
		return
	}

	// 实现grpc服务接口
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, server)

	err = etcdkit.AddRegisterService(logger, etcdkit.Server{
		Name:     ServerName,
		Addr:     ServerAddr,
		Version:  ServerVersion,
		Metadata: "",
	})
	if err != nil {
		logger.Error("注册etcd服务", zap.Error(err))
		return
	}

	logger.Info("启动服务", zap.String("服务名称", ServerName))

	if err := s.Serve(listen); err != nil {
		logger.Error("服务启动失败", zap.Error(err))
		return
	}
}

// shutDown 处理应用关闭
func shutDown() {
	c := make(chan os.Signal, 1)
	// windows 不支持  syscall.SIGTSTP, syscall.SIGSTOP 信号，需要考虑实现
	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	<-c

	fmt.Println("收到关闭信号")
}
