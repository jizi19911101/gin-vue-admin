package v1

import (
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/api_test"
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/autocode"
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/example"
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/project"
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
	ApiTestGroup     api_test.ApiGroup
	ProjectGroup     project.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
