package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/dodoPassport"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/model/request"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/service"
)

type ResourceApi struct{}

var ResourceService = service.ServiceGroupApp.ResourceService

var DoDoUploadImageService = dodoPassport.DoDoGroupApp.DoDoResourcePassport

// UploadImage 上传图片
func (r *ResourceApi) UploadImage(c *gin.Context) {
	_, err := DoDoUploadImageService.UploadResourceToDodo(c)
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("上传失败", c)
		c.Abort()
		return
	}
	response.OkWithMessage("上传成功", c)
}

// CreateResource 创建Resource
func (r *ResourceApi) CreateResource(c *gin.Context) {
	var resource model.DoDoImageSource
	err := c.ShouldBindJSON(&resource)
	if err != nil {
		response.FailWithMessage("创建失败", c)
		c.Abort()
		return
	}
	err = ResourceService.CreateResource(resource)
	if err != nil {
		response.FailWithMessage("创建失败", c)
		c.Abort()
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateResource 修改Resource
func (r *ResourceApi) UpdateResource(c *gin.Context) {
	var resource model.DoDoImageSource
	err := c.ShouldBindJSON(&resource)
	if err != nil {
		response.FailWithMessage("修改失败", c)
		c.Abort()
		return
	}
	err = ResourceService.UpdateResource(resource)
	if err != nil {
		response.FailWithMessage("修改失败", c)
		c.Abort()
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetResourceById 根据ID获取Resource
func (r *ResourceApi) GetResourceById(c *gin.Context) {
	var resource model.DoDoImageSource
	_ = c.ShouldBindJSON(&resource)
	data, err := ResourceService.GetResourceById(resource.ID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		c.Abort()
		return
	}
	response.OkWithData(data, c)
}

// DeleteResource 删除Resource
func (r *ResourceApi) DeleteResource(c *gin.Context) {
	var resource model.DoDoImageSource
	_ = c.ShouldBindJSON(&resource)
	err := ResourceService.DeleteResource(resource.ID)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		c.Abort()
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetResourceList 分页获取Resource列表
func (r *ResourceApi) GetResourceList(c *gin.Context) {
	var pageInfo request.ResourceRequest
	_ = c.ShouldBindJSON(&pageInfo)
	list, total, err := ResourceService.GetResourceList(pageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		c.Abort()
		return
	}
	response.OkWithData(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}
