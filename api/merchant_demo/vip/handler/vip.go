package handler

import (
	"layout-item/api/merchant_demo/vip/facade"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	VipHandler *Server
)

type Server struct {
	service *facade.VipFacade
}

func NewVipHttpHandler() {
	VipHandler = &Server{
		service: facade.NewVipFacade(),
	}
}

// 获取系统当期健康状态
func (s *Server) ServerHealthyNow(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}

// vip获得经验系数配置
func (s *Server) VipExpChangeRatio(ctx *gin.Context) {
	s.service.GetVipConf()
	ctx.String(http.StatusOK, "ok")
}
