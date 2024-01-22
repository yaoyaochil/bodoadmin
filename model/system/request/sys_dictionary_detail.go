package request

import (
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/bodo_request"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	bodo_request.PageInfo
}
