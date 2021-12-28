package autocode

import (
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/api_test"
	"github.com/jizi19911101/gin-vue-admin/server/api/v1/project"
)

type ApiGroup struct {
	// Code generated by github.com/jizi19911101/gin-vue-admin/server Begin; DO NOT EDIT.
	AutoCodeExampleApi
	project.ProjectApi
	api_test.ModuleApi
	project.EnvConfigApi
	//api_test.ApiInfoApi
	// Code generated by github.com/jizi19911101/gin-vue-admin/server End; DO NOT EDIT.
}
