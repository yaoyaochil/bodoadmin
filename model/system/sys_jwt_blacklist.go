package system

import (
	"github.com/yaoyaochil/bodo-admin-server/server/global"
)

type JwtBlacklist struct {
	global.BODO_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
