package request

import (
	"github.com/yaoyaochil/bodo-admin-server/server/model/common/bodo_request"
	"github.com/yaoyaochil/bodo-admin-server/server/plugin/dodo/model"
)

type ResourceRequest struct {
	bodo_request.PageInfo
	model.DoDoImageSource
}
