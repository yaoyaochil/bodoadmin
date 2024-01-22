package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/seazen/service"
)

type MsgApi struct {
}

var sendMsgUserService = service.ServiceGroupApp.SendMsgUserService

// CheckUser 审核用户
func (m *MsgApi) CheckUser(c *gin.Context) {
	var info struct {
		ID uint `json:"id"`
	}
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	err = sendMsgUserService.CheckUser(info.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("审核成功", c)
}
