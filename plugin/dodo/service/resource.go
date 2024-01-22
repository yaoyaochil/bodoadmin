package service

import (
	BODO_Global "github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/model/request"
)

type ResourceService struct{}

// CreateResource 创建Resource
func (r *ResourceService) CreateResource(resource model.DoDoImageSource) (err error) {
	err = BODO_Global.BODO_DB.Create(&resource).Error
	return err
}

// UpdateResource 更新Resource
func (r *ResourceService) UpdateResource(resource model.DoDoImageSource) (err error) {
	err = BODO_Global.BODO_DB.Updates(&resource).Error
	return err
}

// GetResourceById 根据ID获取Resource
func (r *ResourceService) GetResourceById(id uint) (resource model.DoDoImageSource, err error) {
	err = BODO_Global.BODO_DB.Where("id = ?", id).First(&resource).Error
	return resource, err
}

// DeleteResource 删除Resource
func (r *ResourceService) DeleteResource(id uint) (err error) {
	err = BODO_Global.BODO_DB.Where("id = ?", id).Delete(&model.DoDoImageSource{}).Error
	return err
}

// GetResourceList 分页获取Resource列表
func (r *ResourceService) GetResourceList(pageInfo request.ResourceRequest) (resources []model.DoDoImageSource, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := BODO_Global.BODO_DB.Model(&model.DoDoImageSource{})
	var resourceList []model.DoDoImageSource
	if pageInfo.ImageName != "" {
		db = db.Where("image_name LIKE ?", "%"+pageInfo.ImageName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return resourceList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&resourceList).Error
	return resourceList, total, err
}
