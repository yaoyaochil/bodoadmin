package response

import (
	"github.com/yaoyaochil/bodo-admin-server/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
