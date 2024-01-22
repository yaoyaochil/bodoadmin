package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/api"
)

type MenuRouter struct{}

// InitMenuRouter Initialize the menuRouter
func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu")
	var menuApi = api.ApiGroupApp.MenuApi
	{
		//menuRouter.POST("uploadNewsMedia", v1.ApiGroupApp.MediaApi.UploadNewsMedia) // 上传永久图文图片素材
		//menuRouter.POST("uploadOtherMedia", v1.ApiGroupApp.MediaApi.UploadOtherMedia) // 上传其它类型永久素材
		menuRouter.GET("getMenuData", menuApi.GetMenuData) // 获取菜单数据
		menuRouter.POST("createMenu", menuApi.CreateMenu)  // 创建菜单
	}
}
