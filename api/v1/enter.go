package v1

import (
	"github.com/yaoyaochil/bodo-admin-server/server/api/v1/example"
	"github.com/yaoyaochil/bodo-admin-server/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	AutoCodeApi     system.AutoCodeApi
}

var ApiGroupApp = new(ApiGroup)
