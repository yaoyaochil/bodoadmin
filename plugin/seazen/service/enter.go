package service

type ServiceGroup struct {
	SendMsgUserService
	WarnMsgSendService
}

var ServiceGroupApp = new(ServiceGroup)
