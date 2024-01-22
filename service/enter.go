package service

import (
	"github.com/yaoyaochil/bodo-admin-server/server/service/example"
	"github.com/yaoyaochil/bodo-admin-server/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	AutoCodeService     system.AutoCodeService
}

var ServiceGroupApp = new(ServiceGroup)
