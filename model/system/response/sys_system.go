package response

import "github.com/yaoyaochil/bodo-admin-server/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"Config"`
}
