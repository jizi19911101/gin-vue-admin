package router

import (
	"github.com/jizi19911101/gin-vue-admin/server/router/autocode"
	"github.com/jizi19911101/gin-vue-admin/server/router/example"
	"github.com/jizi19911101/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
