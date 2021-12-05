package service

import (
	"github.com/jizi19911101/gin-vue-admin/server/service/api_test"
	"github.com/jizi19911101/gin-vue-admin/server/service/autocode"
	"github.com/jizi19911101/gin-vue-admin/server/service/example"
	"github.com/jizi19911101/gin-vue-admin/server/service/project"
	"github.com/jizi19911101/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
	ApiTestServiceGroup  api_test.ServiceGroup
	GroupServiceGroup    project.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
