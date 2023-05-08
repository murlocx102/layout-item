package etcdkit

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

var etcdAddrs = []string{"127.0.0.1:2379"}

// 根据服务提供grpc连接句柄(用于发现服务)
func GetGRPCConnByService(logger *zap.Logger, app string) (*grpc.ClientConn, error) {
	d := NewServiceDiscovery(etcdAddrs, logger)
	resolver.Register(d)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 设置负载均衡策略：轮循 roundrobin
	conn, err := grpc.DialContext(ctx, BuildResolverUrl(app), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
	cancel()
	if err != nil {
		logger.Error("获取连接句柄失败", zap.Error(err))
		return nil, err
	}

	return conn, nil
}

// 根据服务信息注册etcd(用于服务注册)
func AddRegisterService(logger *zap.Logger, info Server) error {
	r := NewServiceRegister(etcdAddrs, logger)

	// 租约时间默认10s
	_, err := r.Register(info, 10)
	if err != nil {
		logger.Error("etcd注册失败", zap.String("服务名称", info.Name), zap.Error(err))
		return err
	}

	// todo:待完善stop

	return nil
}
