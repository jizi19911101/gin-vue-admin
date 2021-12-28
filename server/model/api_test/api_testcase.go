package api_test

import "github.com/jizi19911101/gin-vue-admin/server/global"

type ApiTestcase struct {
	global.GVA_MODEL
	Name         string
	Api          string
	Module       string
	Organization string
}
