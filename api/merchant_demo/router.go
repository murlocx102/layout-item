package merchant

import (
	"github.com/gin-gonic/gin"
)

// 后台管理接口
func Init(e *gin.Engine) {
	initRouter(e) // 初始化路由模块
}

func initRouter(e *gin.Engine) {
	//异常恢复
	e.Use(gin.Recovery())

	// vip管理模块
	vip := e.Group("/vip")
	{
		initVipRouter(vip)
	}
}
