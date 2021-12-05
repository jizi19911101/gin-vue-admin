package autocode

import (
	"github.com/jizi19911101/gin-vue-admin/server/service/api_test"
	"github.com/jizi19911101/gin-vue-admin/server/service/project"
)

type ServiceGroup struct {
	// Code generated by github.com/jizi19911101/gin-vue-admin/server Begin; DO NOT EDIT.
	AutoCodeExampleService
	project.ProjectService
	api_test.ModuleService
	project.EnvConfigService
	api_test.ApiInfoService
	// Code generated by github.com/jizi19911101/gin-vue-admin/server End; DO NOT EDIT.
}
