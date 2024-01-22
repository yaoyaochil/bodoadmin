package system

import (
	"github.com/yaoyaochil/bodo-admin-server/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"Config"`
}
