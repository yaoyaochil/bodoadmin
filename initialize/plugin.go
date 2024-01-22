package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/middleware"
	xhs "github.com/yaoyaochil/bodo-admin-server/server/plugin/XHS"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/email"
	mqtt_server "github.com/yaoyaochil/bodo-admin-server/server/plugin/mqtt-server"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wecom"
	"github.com/yaoyaochil/bodo-admin-server/server/utils/plugin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权分包模块安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("带鉴权分包模块安装==》", PrivateGroup)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	if global.BODO_DB != nil {
		PluginInit(PrivateGroup, email.CreateEmailPlug(
			global.BODO_CONFIG.Email.To,
			global.BODO_CONFIG.Email.From,
			global.BODO_CONFIG.Email.Host,
			global.BODO_CONFIG.Email.Secret,
			global.BODO_CONFIG.Email.Nickname,
			global.BODO_CONFIG.Email.Port,
			global.BODO_CONFIG.Email.IsSSL,
		))

		// 注册微信插件
		//PluginInit(PublicGroup, wechat.CreateMiniPlug("wxcca489a5c1e8c155", "936da00f8e41d49af9953345ab27d5b8", "bodoQSjSepqkM8QeyqcD", "We0LGGUR1Fmo7tFLrn4LMnjSubkQSjSepqkM8QeyqcD", 99))
		//PluginInit(PublicGroup, wechat.CreateMiniPlug("wx3c031223c0495367", "ee5391fbcabb646632b8293f459d7cc3", "bodoQSjSepqkM8QeyqcD", "We0LGGUR1Fmo7tFLrn4LMnjSubkQSjSepqkM8QeyqcD", 99))
		// 注册DoDo机器人插件
		PluginInit(PublicGroup, dodo.PlugCreateDodo("29218586", "MjkyMTg1ODY.77-9VXPvv70.mhwuYHxcdxTT7mUVtnT4sIfqw0jEb01KpKK8KQBT0yk", "b3ff32b9b8f47ebe77787a04bb5882ad", "176893"))
		// 注册Seazen插件
		//PluginInit(PublicGroup, seazen.CreateSeazenPlug())
		// 注册小红书插件
		PluginInit(PublicGroup, xhs.CreateXHSPlug())

		// 注册MQTT插件
		PluginInit(PublicGroup, mqtt_server.CreateMqttServerPlugin("Api.moonwife.top", "1778", "switch"))

		// 注册企业微信插件
		PluginInit(PublicGroup, wecom.CreateWeComPlug("ww7e16378d3d50ba8e", "EXuOysl2s6qwbJbSdrxmR1coLjhSrxbiaHTWpNlwayE", "1d-0HJf85QYkLw1GwpYJLi-wVuWGxOXJTaWs7PdtAOk", 1000004))
		return
	}
	global.BODO_LOG.Info("未初始化数据库,插件暂不加载")
}
