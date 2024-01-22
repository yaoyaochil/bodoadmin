package xhs

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/XHS/xhs_router"
)

type PluginXHS struct{}

func CreateXHSPlug() *PluginXHS {
	return &PluginXHS{}
}

func (*PluginXHS) Register(group *gin.RouterGroup) {
	xhs_router.RouterGroupApp.InitUserInfoRouter(group)
}

func (*PluginXHS) RouterPath() string {
	return "xhs"
}
