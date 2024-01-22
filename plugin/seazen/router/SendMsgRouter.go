package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/api"
)

type MsgRouter struct {
}

// InitMsgRouter 初始化Msg路由组
func (m *MsgRouter) InitMsgRouter(Router *gin.RouterGroup) {
	msgRouter := Router
	var msgApi = api.ApiGroupApp.MsgApi
	{
		msgRouter.POST("checkUser", msgApi.CheckUser)
	}
}
