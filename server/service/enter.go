package service

import (
	"github/jizi19911101/gin-vue-admin/server/service/autocode"
	"github/jizi19911101/gin-vue-admin/server/service/example"
	"github/jizi19911101/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
