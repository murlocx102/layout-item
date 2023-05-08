//go:build wireinject
// +build wireinject

package service

import (
	"live-service/configs"
	"live-service/internal/vip/repository"

	"github.com/google/wire"
)

func InitUserServer(cfg configs.MYSQL) *VipServer {
	panic(wire.Build(repository.ProviderSet,
		repository.NewVipDbStore,
		wire.Bind(new(repository.VipDbRepoI), new(*repository.VipDbStore)),
		NewVipService,
		NewVipServer))
}
