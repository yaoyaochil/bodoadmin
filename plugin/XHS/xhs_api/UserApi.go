package xhs_api

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/XHS/xhs_passport"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/XHS/xhs_request"
)

type UserApi struct {
}

var UserPassport = xhs_passport.PassportGroupApp.UserInfoPassport

func (u *UserApi) GetUserInfoByID(c *gin.Context) {
	var info xhs_request.UserInfoByIDRequest
	_ = c.ShouldBindJSON(&info)
	data, err := UserPassport.GetUserInfoByUserID(info)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(data.Data, c)
}

func (u *UserApi) GetUserShopList(c *gin.Context) {
	var info xhs_request.GetUserShopRequest
	_ = c.ShouldBindJSON(&info)
	data, err := UserPassport.GetUserShopItemList(info)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(data.Data, c)
}
