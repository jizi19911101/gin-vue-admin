package apicase

import "github.com/jizi19911101/gin-vue-admin/server/global"

type Module struct {
	global.GVA_MODEL
	Name           string
	OrganizationID uint
}

// TableName Module 表名
func (Module) TableName() string {
	return "module"
}
