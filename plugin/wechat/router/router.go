package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/api"
)

type WechatRouter struct {
}

func (s *WechatRouter) InitWechatRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.WechatApi
	getDataApi := api.ApiGroupApp.GetComDataApi
	{
		plugRouter.POST("verifyServer", plugApi.VerifyServer)
		plugRouter.GET("verifyServer", plugApi.VerifyServer)
	}
	{
		plugRouter.POST("getComData", getDataApi.GetComData)
	}
}
