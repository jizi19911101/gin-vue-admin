// 自动生成模板EnvConfig
package project

import (
	"github.com/jizi19911101/gin-vue-admin/server/global"
)

// EnvConfig 结构体
// 如果含有time.Time 请自行import time包
type EnvConfig struct {
	global.GVA_MODEL
	Name     string
	Base_url string
	Project  string
}

// TableName EnvConfig 表名
func (EnvConfig) TableName() string {
	return "env_config"
}
