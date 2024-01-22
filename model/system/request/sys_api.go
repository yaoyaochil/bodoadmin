package request

import (
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/bodo_request"
	"github.com/yaoyaochil/bodo-admin-server/server/model/system"
)

// api分页条件查询及排序结构体
type SearchApiParams struct {
	system.SysApi
	bodo_request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
