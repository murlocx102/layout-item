package test

import (
	"layout-item/configs"
	"layout-item/internal/vip/service"
	"testing"
)

func Test_GetLevel(t *testing.T) {
	err := configs.LoadConfig("../../../configs/", "")
	if err != nil {
		panic("加载配置文件失败" + err.Error())
	}
	cfg := configs.GetConfig()
	server := service.InitUserServer(cfg.MYSQL)
	server.Repo.GetGradeList()
}
