package apicase

import "github.com/jizi19911101/gin-vue-admin/server/global"

type Report struct {
	global.GVA_MODEL
	Name         string
	Url          string
	Organization string
}

// TableName Report 表名
func (Report) TableName() string {
	return "report"
}
