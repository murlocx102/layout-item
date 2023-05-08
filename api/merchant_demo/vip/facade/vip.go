package facade

import (
	"context"
	"fmt"
	"layout-item/internal/vip/service/proto"
	etcdkit "layout-item/pkg/etcd"
	"layout-item/pkg/logger"

	"google.golang.org/grpc"
)

type VipFacade struct {
	vipConn *grpc.ClientConn
}

func NewVipFacade() *VipFacade {
	vip, err := etcdkit.GetGRPCConnByService(logger.Logger, "vip")
	if err != nil {
		panic(err)
	}

	return &VipFacade{
		vipConn: vip,
	}
}

// 业务具体实现聚合
func (v *VipFacade) GetVipConf() {
	c := proto.NewHelloServiceClient(v.vipConn)

	resp, err := c.Hello(context.Background(), &proto.String{Value: "http请求处理"})
	if err != nil {
		fmt.Printf("say hello failed %v", err)

	}

	fmt.Println(resp.GetValue())
}
