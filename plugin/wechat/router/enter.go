package router

type RouterGroup struct {
	WechatRouter
	WxLoginRouter
	SendMsgROuter
	QrRouter
	MediaRouter
	MenuRouter
}

var RouterGroupApp = new(RouterGroup)
