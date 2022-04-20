package monkey

import "github.com/jizi19911101/gin-vue-admin/server/global"

type Task struct {
	global.GVA_MODEL
	Device   string
	Duration string
	App      string
	Report   string
	CleanLog *bool
	UserId   string
	Status   string
	//OrganizationID uint
}

// TableName Report 表名
func (Task) TableName() string {
	return "task"
}
