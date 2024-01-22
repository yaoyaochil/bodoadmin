package xhs_router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/XHS/xhs_api"
)

type UserRouter struct{}

func (u *UserRouter) InitUserInfoRouter(Router *gin.RouterGroup) {
	pluginRouter := Router
	pluginApi := xhs_api.ApiGroupApp.UserApi
	{
		pluginRouter.POST("getUserInfoByXhsId", pluginApi.GetUserInfoByID)
		pluginRouter.POST("getUserShopList", pluginApi.GetUserShopList)
	}
}
