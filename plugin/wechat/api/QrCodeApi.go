package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/Wxglobal"
)

type QrCodeApi struct {
}

type QrCodeCreateRequest struct {
	Scene string `json:"scene"`
}

func (q *QrCodeApi) CreateForeverQrCode(c *gin.Context) {
	var scene QrCodeCreateRequest
	err := c.ShouldBindJSON(&scene)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	forever, err := Wxglobal.GlobalConfig.OfficialAccountApp.QRCode.Forever(Wxglobal.GlobalConfig.NullCtx, scene.Scene)
	if err != nil {
		return
	}
	response.OkWithData(forever, c)
}
