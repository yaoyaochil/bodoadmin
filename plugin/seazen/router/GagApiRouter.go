package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/api"
)

type GagRouter struct {
}

func (g *GagRouter) InitGagRouter(Router *gin.RouterGroup) {
	gagRouter := Router
	var gagApi = api.ApiGroupApp.GagApi
	{
		gagRouter.POST("getShopEntityCheck", gagApi.GetShopEntityCheckApi)
		gagRouter.POST("testSendMsg", gagApi.TestSendMsgApi)
	}
	{
		gagRouter.POST("getWarnInfoDataByID", gagApi.GetWarnInfoDataByID)
	}
}
