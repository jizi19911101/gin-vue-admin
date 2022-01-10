package request

import "github.com/jizi19911101/gin-vue-admin/server/model/common/request"

type ApiSearch struct {
	request.PageInfo
	Name           string `json:"name" form:"name"`
	Module         string `json:"module" form:"module"`
	OrganizationID uint   `json:"organizationID" form:"organizationID"`
}
