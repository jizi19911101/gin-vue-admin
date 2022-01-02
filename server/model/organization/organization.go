// 自动生成模板Project
package organization

import (
	"github.com/jizi19911101/gin-vue-admin/server/global"
)

// Project 结构体
// 如果含有time.Time 请自行import time包
type Organization struct {
	global.GVA_MODEL
	Name string
}

// TableName Project 表名
func (Organization) TableName() string {
	return "organization"
}
