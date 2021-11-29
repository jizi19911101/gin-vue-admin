// 自动生成模板ApiInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ApiInfo 结构体
// 如果含有time.Time 请自行import time包
type ApiInfo struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:接口名称;type:varchar;"`
      Method  string `json:"method" form:"method" gorm:"column:method;comment:请求方法;type:char;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:请求url;type:varchar;"`
      Params  string `json:"params" form:"params" gorm:"column:params;comment:请求参数;type:varchar;"`
      Project  string `json:"project" form:"project" gorm:"column:project;comment:所属项目;type:char;"`
      Module  string `json:"module" form:"module" gorm:"column:module;comment:所属模块;type:varchar;"`
}


// TableName ApiInfo 表名
func (ApiInfo) TableName() string {
  return "apiInfo"
}

