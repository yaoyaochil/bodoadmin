package api

import (
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/menu/request"
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/Wxglobal"
)

type MenuApi struct {
}

// GetMenuData 获取菜单数据
func (m *MenuApi) GetMenuData(c *gin.Context) {
	data, err := Wxglobal.GlobalConfig.OfficialAccountApp.Menu.Get(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

// CreateMenu 创建菜单
func (m *MenuApi) CreateMenu(c *gin.Context) {
	var menu []*request.Button

	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("menu:", menu)
	data, err := Wxglobal.GlobalConfig.OfficialAccountApp.Menu.Create(c, menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if data.ErrCode != 0 {
		response.FailWithMessage(data.ErrMsg, c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
