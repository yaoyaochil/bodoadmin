package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/api"
)

type WxLoginRouter struct{}

func (s *WxLoginRouter) InitWxLoginRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.WxLoginApi
	{
		plugRouter.GET("getLoginPic", plugApi.GetLoginPic)
		plugRouter.GET("loginOrCreate", plugApi.LoginOrCreate)
		plugRouter.POST("clearWx", plugApi.ClearWx)
		// 检查是否绑定
		plugRouter.GET("checkBind", plugApi.CheckBind)
	}
}
