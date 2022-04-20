package service

import (
	"github.com/jizi19911101/gin-vue-admin/server/service/apicase"
	"github.com/jizi19911101/gin-vue-admin/server/service/autocode"
	"github.com/jizi19911101/gin-vue-admin/server/service/example"
	"github.com/jizi19911101/gin-vue-admin/server/service/monkey"
	"github.com/jizi19911101/gin-vue-admin/server/service/organization"
	"github.com/jizi19911101/gin-vue-admin/server/service/sync"
	"github.com/jizi19911101/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	ExampleServiceGroup      example.ServiceGroup
	AutoCodeServiceGroup     autocode.ServiceGroup
	SyncServiceGroup         sync.ServiceGroup
	OrganizationServiceGroup organization.ServiceGroup
	ApiCaseServiceGroup      apicase.ServiceGroup
	MonkeyServiceGroup       monkey.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
