package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/api"
)

type MediaRouter struct{}

func (m *MediaRouter) InitMediaRouter(Router *gin.RouterGroup) {
	mediaRouter := Router.Group("media")
	var mediaApi = api.ApiGroupApp.MediaApi
	{
		mediaRouter.POST("uploadNewsMedia", mediaApi.UploadNewsMedia)
		mediaRouter.POST("uploadOtherMedia", mediaApi.UploadOtherMedia)
		mediaRouter.POST("getMediaList", mediaApi.GetMediaList)
		mediaRouter.POST("getMediaByID", mediaApi.GetMediaById)
		mediaRouter.POST("deleteMediaByID", mediaApi.DeleteMediaById)
	}
}
