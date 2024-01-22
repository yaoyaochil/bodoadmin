package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/api"
)

type SeazenRouter struct{}

func (s *SeazenRouter) InitSeazenRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.QrCodeApi
	{
		plugRouter.POST("createQrCode", plugApi.CreateForeverQrCode)
	}
}
