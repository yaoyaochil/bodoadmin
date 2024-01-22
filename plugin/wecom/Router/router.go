package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wecom/Api"
)

type WecomRouter struct {
}

func (s *WecomRouter) InitWeComRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := Api.ApiGroupApp.CallBackApi
	{
		plugRouter.GET("/callback", plugApi.CallBack)
	}
	{
		plugRouter.POST("/callback", plugApi.CallBack)
	}
}
