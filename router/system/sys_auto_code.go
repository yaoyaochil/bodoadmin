package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yaoyaochil/bodo-admin-server/server/api/v1"
)

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autoCode")
	autoCodeApi := v1.ApiGroupApp.SystemApiGroup.AutoCodeApi
	{
		autoCodeRouter.POST("createPlug", autoCodeApi.AutoPlug) // 自动插件包模板
	}
}
