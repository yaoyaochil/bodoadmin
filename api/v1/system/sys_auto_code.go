package system

import (
	"strings"

	"github.com/yaoyaochil/bodo-admin-server/server/global"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AutoCodeApi struct{}

var caser = cases.Title(language.English)

// AutoPlug
// @Tags      AutoCode
// @Summary   创建插件模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysAutoCode                                         true  "创建插件模板"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "创建插件模板成功"
// @Router    /autoCode/createPlug [post]
func (autoApi *AutoCodeApi) AutoPlug(c *gin.Context) {
	var a system.AutoPlugReq
	err := c.ShouldBindJSON(&a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	a.Snake = strings.ToLower(a.PlugName)
	a.NeedModel = a.HasRequest || a.HasResponse
	err = autoCodeService.CreatePlug(a)
	if err != nil {
		global.BODO_LOG.Error("预览失败!", zap.Error(err))
		response.FailWithMessage("预览失败", c)
		return
	}
	response.Ok(c)
}
