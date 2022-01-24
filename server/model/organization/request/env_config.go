package request

import (
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
)

type EnvConfigSearch struct {
	Name         string `json:"name" form:"name"  validate:"required"`
	Base_url     string `json:"base_url" form:"base_url"   validate:"required"`
	Organization string `json:"organization" form:"organization" gorm:"column:organization;comment:"  validate:"required"`
	request.PageInfo
}

type EnvConfigReq struct {
	ID           uint
	Name         string `json:"name" form:"name"  validate:"required"`
	Base_url     string `json:"base_url" form:"base_url"   validate:"required"`
	Organization string `json:"organization" form:"organization"   validate:"required"`
}
