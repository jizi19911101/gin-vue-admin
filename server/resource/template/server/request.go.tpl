package request

import (
	"github.com/jizi19911101/gin-vue-admin/server/model/autocode"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}