package router

type RouterGroup struct {
	SeazenRouter
	GagRouter
	MsgRouter
}

var RouterGroupApp = new(RouterGroup)
