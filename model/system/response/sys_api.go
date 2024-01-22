package response

import "github.com/yaoyaochil/bodo-admin-server/server/model/system"

type SysAPIResponse struct {
	Api system.SysApi `json:"xhs_api"`
}

type SysAPIListResponse struct {
	Apis []system.SysApi `json:"apis"`
}
