package monkey

import "github.com/jizi19911101/gin-vue-admin/server/global"

type MonkeyReport struct {
	global.GVA_MODEL
	Name    string
	Content string ` gorm:"type:mediumtext"`
	//OrganizationID uint
}

// TableName Report 表名
func (MonkeyReport) TableName() string {
	return "monkey_report"
}
