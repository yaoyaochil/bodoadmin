package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/service"
)

var GetComDataService = service.ServiceGroupApp.ComDataService

type GetComDataApi struct {
}

// GetComData 获取企业微信数据
func (g *GetComDataApi) GetComData(c *gin.Context) {
	var info service.Dates
	_ = c.ShouldBindJSON(&info)
	UserSummary, err := GetComDataService.GetComData(info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	response.OkWithData(UserSummary, c)
}
