package Wxglobal

import (
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/wechat/config"
	"sync"
)

var GlobalConfig = new(config.WeComConfig)

type CodeInfo struct {
	CanLogin   bool
	OpenId     string
	CreateTime int64
}

var LoginCodeMap = sync.Map{}

var BindMap = sync.Map{}

var SUBSCRIBE = "subscribe"     //(订阅 关注)
var UNSUBSCRIBE = "unsubscribe" //(退订 取消关注)
var SCAN = "SCAN"               //(扫码)
