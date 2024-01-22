package api

type ApiGroup struct {
	QrCodeApi
	GagApi
	MsgApi
}

var ApiGroupApp = new(ApiGroup)
