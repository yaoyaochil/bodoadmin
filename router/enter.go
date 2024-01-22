package router

import (
	"github.com/yaoyaochil/bodo-admin-server/server/router/example"
	"github.com/yaoyaochil/bodo-admin-server/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
