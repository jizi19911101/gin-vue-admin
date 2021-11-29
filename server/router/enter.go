package router

import (
	"github/jizi19911101/gin-vue-admin/server/router/autocode"
	"github/jizi19911101/gin-vue-admin/server/router/example"
	"github/jizi19911101/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
