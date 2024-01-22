package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/model"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/service"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/wx_request"
	"github.com/yaoyaochil/bodo-admin-server/server/utils"
)

type SendMsgApi struct {
}

var SendMsgService = service.ServiceGroupApp.WxMsgService

func (s *SendMsgApi) SendMsg(c *gin.Context) {
	var msg model.WxMsg
	_ = c.ShouldBindJSON(&msg)
	if err := utils.Verify(msg, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := SendMsgService.SendMsg(msg); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("发送成功", c)
	}
}

// SendTemplateMsg 发送模版消息
func (s *SendMsgApi) SendTemplateMsg(c *gin.Context) {
	var msg wx_request.TemplateMessageRequest
	_ = c.ShouldBindJSON(&msg)
	if data, err := SendMsgService.SendWarningTemplateMsg(msg); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(data, c)
	}
}
