package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wecom/Api"
)

type DepartmentRouter struct{}

// 初始化路由
func (departmentRouter *DepartmentRouter) InitDepartmentRouter(Router *gin.RouterGroup) {
	departmentRouterApi := Router.Group("department")
	departmentApi := Api.ApiGroupApp.DepartmentApi
	{
		departmentRouterApi.GET("getDepartmentList", departmentApi.GetDepartmentList) // 获取部门列表
		departmentRouterApi.POST("createDepartment", departmentApi.CreateDepartment)  // 创建部门
		departmentRouterApi.POST("updateDepartment", departmentApi.UpdateDepartment)  // 更新部门
		departmentRouterApi.POST("deleteDepartment", departmentApi.DeleteDepartment)  // 删除部门
	}
}
