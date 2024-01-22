package service

import (
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/Wxglobal"
)

type ComDataService struct{}

type Dates struct {
	StartDate string `json:"start_date"` //格式为2020-08-01
	EndDate   string `json:"end_date"`   // 格式为2020-08-07
}

// GetComData 获取获取用户增减数据
func (e *ComDataService) GetComData(date Dates) (data interface{}, err error) {
	// 获取7天前的日期

	if date.StartDate == "" || date.EndDate == "" {
		date.StartDate = "2020-08-01"
		date.EndDate = "2020-08-07"
	}
	// 获取用户增减数据
	UserCumulate, _ := Wxglobal.GlobalConfig.OfficialAccountApp.DataCube.GetUserSummary(Wxglobal.GlobalConfig.NullCtx, date.StartDate, date.EndDate)
	return UserCumulate, err
}
