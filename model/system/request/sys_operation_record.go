package request

import (
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/bodo_request"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	bodo_request.PageInfo
}
