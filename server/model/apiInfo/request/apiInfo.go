package request

import "github/jizi19911101/gin-vue-admin/server/global"

type ApiInfo struct {
	global.GVA_MODEL
	Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar;"`
	Method  string `json:"method" form:"method" gorm:"column:method;comment:;type:char;"`
	Url  string `json:"url" form:"url" gorm:"column:url;comment:;type:varchar;"`
	Params  []string `json:"params" form:"params" `
	Project  string `json:"project" form:"project" gorm:"column:project;comment:;type:char;"`
	Module  string `json:"module" form:"module" gorm:"column:module;comment:;type:varchar;"`
}
