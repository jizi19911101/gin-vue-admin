package api_test

import "github.com/jizi19911101/gin-vue-admin/server/global"

type ModuleInfo struct {
	global.GVA_MODEL
	Name         string
	Organization string
}
