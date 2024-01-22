package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/api"
)

type SendMsgROuter struct{}

func (s *SendMsgROuter) InitSendMsgRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.SendMsgApi
	{
		plugRouter.POST("sendMsg", plugApi.SendMsg)
		plugRouter.POST("sendTemplateMsg", plugApi.SendTemplateMsg)
	}
}
