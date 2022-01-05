package apiTest

import "github.com/jizi19911101/gin-vue-admin/server/global"

type ApiTestcase struct {
	global.GVA_MODEL
	Name           string
	Class          string
	Api            string
	Module         string
	OrganizationID uint
}

// TableName ApiTestcase 表名
func (ApiTestcase) TableName() string {
	return "api_testcase"
}
