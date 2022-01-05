package apiTest

import (
	"github.com/jizi19911101/gin-vue-admin/server/global"
)

type Api struct {
	global.GVA_MODEL
	Name           string
	Module         string
	OrganizationID uint
}

// TableName Api 表名
func (Api) TableName() string {
	return "api"
}
