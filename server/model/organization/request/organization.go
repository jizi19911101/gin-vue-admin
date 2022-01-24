package request

import (
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
)

type OrganizationSearch struct {
	Name string `json:"name" form:"name"  validate:"required"`
	request.PageInfo
}

type OrganizationReq struct {
	ID   uint
	Name string `json:"name" form:"name"  validate:"required"`
}
