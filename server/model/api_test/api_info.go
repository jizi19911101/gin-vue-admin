// 自动生成模板ApiInfo
package api_test

import (
	"github.com/jizi19911101/gin-vue-admin/server/global"
)

// ApiInfo 结构体
// 如果含有time.Time 请自行import time包
type ApiInfo struct {
	global.GVA_MODEL
	Name    string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar;"`
	Method  string `json:"method" form:"method" gorm:"column:method;comment:;type:char;"`
	Url     string `json:"url" form:"url" gorm:"column:url;comment:;type:varchar;"`
	Params  string `json:"params" form:"params" gorm:"column:params;comment:;type:varchar;"`
	Project string `json:"project" form:"project" gorm:"column:project;comment:;type:char;"`
	Module  string `json:"module" form:"module" gorm:"column:module;comment:;type:varchar;"`
}

// TableName ApiInfo 表名
func (ApiInfo) TableName() string {
	return "apiInfo"
}
