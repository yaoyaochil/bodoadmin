package Api

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/department/request"
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/response"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wecom/ComGlobal"
)

type DepartmentApi struct{}

/**
 * @Description: 获取部门列表
 * @param id 部门id
 * @return {Department} 部门列表数据
 */
func (d *DepartmentApi) GetDepartmentList(c *gin.Context) {
	var id int
	err := c.ShouldBindQuery(&id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := ComGlobal.GlobalConfig.WeComApp.Department.List(c, id)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(data, c)
}

/**
 * @Description: 创建部门
 * @param c
 */
func (d *DepartmentApi) CreateDepartment(c *gin.Context) {
	var request *request.RequestDepartmentInsert

	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := ComGlobal.GlobalConfig.WeComContactApp.Department.Create(c, request)
	if err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithData(data, c)
}

/**
 * @Description: 更新部门
 * @param c
 */
func (d *DepartmentApi) UpdateDepartment(c *gin.Context) {
	var request *request.RequestDepartmentUpdate

	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := ComGlobal.GlobalConfig.WeComContactApp.Department.Update(c, request)
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithData(data, c)
}

/**
 * @Description: 删除部门
 * @param c
 */
func (d *DepartmentApi) DeleteDepartment(c *gin.Context) {
	var id int
	err := c.ShouldBindQuery(&id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	_, err = ComGlobal.GlobalConfig.WeComContactApp.Department.Delete(c, id)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok(c)
}
