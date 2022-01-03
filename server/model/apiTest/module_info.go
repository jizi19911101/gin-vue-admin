package apiTest

import "github.com/jizi19911101/gin-vue-admin/server/global"

type ModuleInfo struct {
	global.GVA_MODEL
	Name         string
	Organization string
}

// TableName ModuleInfo 表名
func (ModuleInfo) TableName() string {
	return "module_info"
}
