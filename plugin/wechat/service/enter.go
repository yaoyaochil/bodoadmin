package service

type ServiceGroup struct {
	WechatService
	ComDataService
	WxMsgService
	WxLoginService
}

var ServiceGroupApp = new(ServiceGroup)
