package request

import (
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/organization"
)

type OrganizationSearch struct {
	organization.Organization
	request.PageInfo
}

type OrganizationRes struct {
	ID   uint
	Name string `json:"name" form:"name" gorm:"column:name;comment:" validate:"required"`
}
