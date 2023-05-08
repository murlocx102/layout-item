package pool

import (
	"errors"
	"layout-item/pkg/errkit"
	"time"
)

var (
	GrpcErr     = errkit.New(300, errors.New("grpc err"))
	GrpcCallErr = errkit.New(301, errors.New("grpc call err"))
)

// grpc连接池配置
type PoolOption struct {
	MinSize        int
	MaxSize        int
	IdleTimeout    time.Duration
	MaxLifeTimeout time.Duration
	GetConnTimeout time.Duration //获取连接超时时间
}

// 客户端配置选项
type ClientOption struct {
	*PoolOption
	Server       string
	WithBlock    bool
	WithInsecure bool
}
