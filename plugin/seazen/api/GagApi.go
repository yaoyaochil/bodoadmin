package api

import (
	"github.com/gin-gonic/gin"
	BODO_Global "github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/passport"
	seazenService "github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/service"
)

type GagApi struct {
}

var GagPassport = passport.PassportGroupApp.GagService
var WarnServer = seazenService.ServiceGroupApp.WarnMsgSendService

func (g *GagApi) GetShopEntityCheckApi(c *gin.Context) {
	data, err := GagPassport.GetShopEntityCheck()
	if err != nil {
		response.FailWithMessage("获取失败", c)
	}
	response.OkWithData(data, c)
}

func (g *GagApi) GetShopEntitySaleInfoApi(c *gin.Context) {
	data, err := GagPassport.GetShopEntitySaleInfo()
	if err != nil {
		response.FailWithMessage("获取失败", c)
	}
	response.OkWithData(data, c)
}

func (g *GagApi) TestSendMsgApi(c *gin.Context) {
	err := WarnServer.SendWarningTemplateMsg()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("发送成功", c)
}

func (g *GagApi) GetWarnInfoDataByID(c *gin.Context) {
	var info model.WarnMsgInfoList
	err := c.ShouldBindJSON(&info)
	var data model.WarnMsgInfoList
	err = BODO_Global.BODO_DB.Model(model.WarnMsgInfoList{}).Where("id = ?", info.ID).First(&data).Error
	if err != nil {
		response.FailWithMessage("无数据", c)
		c.Abort()
		return
	}
	response.OkWithData(data, c)
}
