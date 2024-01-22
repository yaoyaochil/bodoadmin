package Api

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/department/request"
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wecom/ComGlobal"
)

type DepartmentUserApi struct {
}

/**
 * @Description: 获取部门成员
 * @param c
 */
func (d *DepartmentUserApi) GetDepartmentUserList(c *gin.Context) {
	var request *request.RequestDepartmentUserList

	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := ComGlobal.GlobalConfig.WeComContactApp.User.GetDepartmentUsers(c, request)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(data, c)
}
