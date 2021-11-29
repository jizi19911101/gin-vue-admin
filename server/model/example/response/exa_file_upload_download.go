package response

import "github.com/jizi19911101/gin-vue-admin/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
