package api_test

import (
	"github.com/jizi19911101/gin-vue-admin/server/global"
)

type ApiInfo struct {
	global.GVA_MODEL
	Name         string
	Module       string
	Organization string
}
