package request

import (
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/bodo_request"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	bodo_request.PageInfo
}
