// 自动生成模板Project
package project

import (
	"github.com/jizi19911101/gin-vue-admin/server/global"
)

// Project 结构体
// 如果含有time.Time 请自行import time包
type Project struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:项目名称"`
}

// TableName Project 表名
func (Project) TableName() string {
	return "project"
}
