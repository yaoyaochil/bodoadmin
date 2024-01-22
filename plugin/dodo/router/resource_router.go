package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/api"
)

type ResourceRouter struct {
}

// InitResourceRouter 注册路由
func (r *ResourceRouter) InitResourceRouter(Router *gin.RouterGroup) {
	resourceRouter := Router.Use()
	plugApi := api.ApiGroupApp.ResourceApi
	{
		resourceRouter.POST("createResource", plugApi.CreateResource)
		resourceRouter.POST("updateResource", plugApi.UpdateResource)
		resourceRouter.POST("deleteResource", plugApi.DeleteResource)
	}
	{
		resourceRouter.POST("getResourceList", plugApi.GetResourceList)
		resourceRouter.POST("getResourceById", plugApi.GetResourceById)
	}
	{
		resourceRouter.POST("uploadImage", plugApi.UploadImage)
	}
}
