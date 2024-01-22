package model

import BODO_Global "github.com/yaoyaochil/bodo-admin-server/server/global"

type SeazenConfig struct {
	BODO_Global.BODO_MODEL
	Cookie string `gorm:"column:cookie;type:long;" json:"cookie"`
	Name   string `gorm:"column:name;type:varchar(255);not null;unique" json:"name"`
	Desc   string `gorm:"column:desc;type:varchar(255);not null;unique" json:"desc"`
}

type SendMsgUserList struct {
	BODO_Global.BODO_MODEL
	OpenID    string `gorm:"column:open_id;type:varchar(255);not null;unique" json:"open_id"`
	IsEnabled bool   `gorm:"column:is_enabled;type:tinyint(1);not null;default:1" json:"is_enabled"`
}

type SendMsgType struct {
	BODO_Global.BODO_MODEL
	Name string `gorm:"column:name;type:varchar(255);not null;unique;comment:消息用途" json:"name"`
	Desc string `gorm:"column:desc;type:varchar(255);not null;unique;comment:描述" json:"desc"`
}

// QrCodeList 永久二维码列表
type QrCodeList struct {
	BODO_Global.BODO_MODEL
	Name     string `gorm:"column:name;type:varchar(255);not null;unique;comment:二维码名称" json:"name"`
	Desc     string `gorm:"column:desc;type:varchar(255);not null;unique;comment:描述" json:"desc"`
	ImageUrl string `gorm:"column:image_url;type:varchar(255);not null;unique;comment:二维码地址" json:"image_url"`
}

type WarnMsgInfoList struct {
	BODO_Global.BODO_MODEL
	SendDate      string `gorm:"column:send_date;type:varchar(255);not null;unique;comment:发送日期" json:"send_date"`
	OpenID        string `gorm:"column:open_id;type:long;comment:用户open_id" json:"open_id"`
	NotDeviceShop string `gorm:"column:not_device_shop;type:long;comment:未绑定设备的店铺" json:"not_device_shop"`
	HaveDiffAbs   string `gorm:"column:have_diff_abs;type:long;comment:有差异的店铺" json:"have_diff_abs"`
}
