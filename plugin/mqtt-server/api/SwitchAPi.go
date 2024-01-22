package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/mqtt-server/service"
)

type SwitchApi struct {
}

var Service = service.ServiceGroupApp.SwitchService

func (s *SwitchApi) SwitchOn(c *gin.Context) {
	if err := Service.SwitchOn(); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithMessage("开关打开成功！", c)
}

func (s *SwitchApi) SwitchOff(c *gin.Context) {
	if err := Service.SwitchOff(); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithMessage("开关关闭成功！", c)
}
