package service

import (
	"context"
	"fmt"
	"layout-item/internal/vip/service/proto"
	"strings"
)

type VipServer struct {
	Repo *VipService
}

func NewVipServer(service *VipService) *VipServer {
	return &VipServer{
		Repo: service,
	}
}

func (s *VipServer) Hello(ctx context.Context, in *proto.String) (*proto.String, error) {
	fmt.Println("已调用Hello服务")
	result := strings.Join([]string{in.Value, "hello world"}, ",")
	return &proto.String{
		Value: result,
	}, nil
}
