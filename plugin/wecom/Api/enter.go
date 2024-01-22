package Api

type ApiGroup struct {
	CallBackApi
	DepartmentApi
}

var ApiGroupApp = new(ApiGroup)
