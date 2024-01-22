package Router

type RouterGroup struct {
	WecomRouter
	DepartmentRouter
}

var RouterGroupApp = new(RouterGroup)
