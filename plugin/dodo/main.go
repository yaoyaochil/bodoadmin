package dodo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	BODO_Global "github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/dodo_global"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/router"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/plugin-tool/utils"
)

type PluginDodo struct {
}

func PlugCreateDodo(ClientID string, Token string, Secret string, IslandSourceId string) *PluginDodo {
	dodo_global.GlobalConfig.ClientID = ClientID
	dodo_global.GlobalConfig.Token = Token
	dodo_global.GlobalConfig.Secret = Secret
	dodo_global.GlobalConfig.Authorization = fmt.Sprintf("Bot %s.%s", dodo_global.GlobalConfig.ClientID, dodo_global.GlobalConfig.Token)
	dodo_global.GlobalConfig.IslandSourceId = IslandSourceId
	err := BODO_Global.BODO_DB.AutoMigrate(
		&model.DoDoImageSource{},
	)
	if err != nil {
		panic(err)
	}
	// 注册API
	utils.RegisterApis(
		system.SysApi{
			Path:        "/dodo/createResource",
			Description: "创建资源",
			ApiGroup:    "DoDo资源中心",
			Method:      "POST",
		},
		system.SysApi{
			Path:        "/dodo/updateResource",
			Description: "更新资源",
			ApiGroup:    "DoDo资源中心",
			Method:      "POST",
		},
		system.SysApi{
			Path:        "/dodo/deleteResource",
			Description: "删除资源",
			ApiGroup:    "DoDo资源中心",
			Method:      "POST",
		},
		system.SysApi{
			Path:        "/dodo/getResourceList",
			Description: "获取资源列表",
			ApiGroup:    "DoDo资源中心",
			Method:      "POST",
		},
		system.SysApi{
			Path:        "/dodo/getResourceById",
			Description: "根据ID获取资源",
			ApiGroup:    "DoDo资源中心",
			Method:      "POST",
		},
		system.SysApi{
			Path:        "/dodo/uploadImage",
			Description: "上传图片",
			ApiGroup:    "DoDo资源中心",
			Method:      "POST",
		},
	)
	return &PluginDodo{}
}

func (*PluginDodo) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitResourceRouter(group)
}

func (*PluginDodo) RouterPath() string {
	return "dodo"
}
