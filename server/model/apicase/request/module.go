package request

import "github.com/jizi19911101/gin-vue-admin/server/model/common/request"

type ModuleSearch struct {
	request.PageInfo
	Name           string `json:"name" form:"name" `
	organizationID uint   `json:"organizationID" form:"organizationID" `
}
