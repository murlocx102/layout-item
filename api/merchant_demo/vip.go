package merchant

import (
	"layout-item/api/merchant_demo/vip/handler"

	"github.com/gin-gonic/gin"
)

func initVipRouter(m *gin.RouterGroup) {
	{
		// vip获得经验系数配置
		m.POST("/exp/change/ratio", handler.VipHandler.VipExpChangeRatio)

		m.GET("/healthy", handler.VipHandler.ServerHealthyNow) // 服务健康检查接口,正常返回ok
	}
}
