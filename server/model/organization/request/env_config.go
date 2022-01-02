package request

import (
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/organization"
)

type EnvConfigSearch struct {
	organization.EnvConfig
	request.PageInfo
}

type EnvConfigReq struct {
	ID           uint
	Name         string `json:"name" form:"name" gorm:"column:name;comment:" validate:"required"`
	Base_url     string `json:"base_url" form:"base_url" gorm:"column:base_url;comment:"  validate:"required"`
	Organization string `json:"organization" form:"organization" gorm:"column:organization;comment:"  validate:"required"`
}
