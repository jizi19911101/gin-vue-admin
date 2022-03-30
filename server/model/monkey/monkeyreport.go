package monkey

import "github.com/jizi19911101/gin-vue-admin/server/global"

type MonkeyReport struct {
	global.GVA_MODEL
	Name         string
	AppName      string
	AppVersion   string
	Duration     string
	BeginTime    string
	PhoneSystem  string
	PhoneVersion string
	Log          string ` gorm:"type:text"`
	//OrganizationID uint
}

// TableName Report 表名
func (MonkeyReport) TableName() string {
	return "monkey_report"
}
