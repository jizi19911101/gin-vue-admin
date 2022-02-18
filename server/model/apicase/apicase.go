package apicase

import "github.com/jizi19911101/gin-vue-admin/server/global"

type ApiCase struct {
	global.GVA_MODEL
	Name           string
	Class          string
	Api            string
	Module         string
	Title          string
	OrganizationID uint
}

// TableName ApiTestcase 表名
func (ApiCase) TableName() string {
	return "apicase"
}
