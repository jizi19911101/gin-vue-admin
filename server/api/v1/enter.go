package v1

import (
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/autocode"
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/example"
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/organization"
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/sync"
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	AutoCodeApiGroup  autocode.ApiGroup
	SyncGroup         sync.ApiGroup
	OrganizationGroup organization.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
