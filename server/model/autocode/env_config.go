// 自动生成模板EnvConfig
package autocode

import (
	"github/jizi19911101/gin-vue-admin/server/global"
)

// EnvConfig 结构体
// 如果含有time.Time 请自行import time包
type EnvConfig struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:"`
      Base_url  string `json:"base_url" form:"base_url" gorm:"column:base_url;comment:"`
      Project  string `json:"project" form:"project" gorm:"column:project;comment:"`

}


// TableName EnvConfig 表名
func (EnvConfig) TableName() string {
  return "env_config"
}

