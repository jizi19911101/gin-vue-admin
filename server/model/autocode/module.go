// 自动生成模板Module
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Module 结构体
// 如果含有time.Time 请自行import time包
type Module struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:"`
      Project  string `json:"project" form:"project" gorm:"column:project;comment:"`
}


// TableName Module 表名
func (Module) TableName() string {
  return "module"
}

