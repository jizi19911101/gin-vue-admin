package request

import "github.com/jizi19911101/gin-vue-admin/server/model/common/request"

type RunApiCaseReq struct {
	OrganizationID uint   `json:"organizationID" form:"organizationID" `
	Module         string `json:"module" form:"module"  `
	Api            string `json:"api" form:"api"  `
	Case           string `json:"case" form:"case"  `
	Env            string `json:"env" form:"env"  validate:"required"`
}

type ApiCaseSearch struct {
	request.PageInfo
	OrganizationID uint   `json:"organizationID" form:"organizationID" `
	Module         string `json:"module" form:"module"  `
	Api            string `json:"api" form:"api"  `
	Name           string `json:"name" form:"name"  `
}
