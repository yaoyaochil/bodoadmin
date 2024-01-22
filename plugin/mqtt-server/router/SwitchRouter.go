package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/mqtt-server/api"
)

type SwitchRouter struct {
}

// InitSwitchRouterGroup 注册路由
func (s *SwitchRouter) InitSwitchRouterGroup(router *gin.RouterGroup) {
	switchRouter := router.Group("switch")
	plugApi := api.ApiGroupApp.SwitchApi
	{
		switchRouter.POST("on", plugApi.SwitchOn)
		switchRouter.POST("off", plugApi.SwitchOff)
	}
}
