package initialize

import (
	_ "github.com/yaoyaochil/bodo-admin-server/server/source/example"
	_ "github.com/yaoyaochil/bodo-admin-server/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
