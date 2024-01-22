package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/api"
)

type QrRouter struct{}

func (s *QrRouter) InitQrRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.QrCodeApi
	{
		plugRouter.POST("createQrCode", plugApi.CreateForeverQrCode)
	}
}
