package api

type ApiGroup struct {
	WechatApi
	GetComDataApi
	WxLoginApi
	SendMsgApi
	QrCodeApi
	MediaApi
	MenuApi
}

var ApiGroupApp = new(ApiGroup)
