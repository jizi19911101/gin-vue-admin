package v1

import (
	"github/jizi19911101/gin-vue-admin/server/api/v1/autocode"
	"github/jizi19911101/gin-vue-admin/server/api/v1/example"
	"github/jizi19911101/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
