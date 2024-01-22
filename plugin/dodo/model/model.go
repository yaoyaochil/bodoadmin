package model

import (
	BODO_Global "github.com/yaoyaochil/bodo-admin-server/server/global"
)

type DoDoImageSource struct {
	BODO_Global.BODO_MODEL
	ImageName string `json:"imageName" gorm:"column:image_name;comment:图片名称"`
	Height    int    `json:"height" gorm:"column:height;comment:图片高度"`
	Url       string `json:"url" gorm:"column:url;comment:图片地址"`
	Width     int    `json:"width" gorm:"column:width;comment:图片宽度"`
}

func (DoDoImageSource) TableName() string {
	return "dodo_image_source"
}
