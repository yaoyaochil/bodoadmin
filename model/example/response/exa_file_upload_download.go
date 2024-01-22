package response

import "github.com/yaoyaochil/bodo-admin-server/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
