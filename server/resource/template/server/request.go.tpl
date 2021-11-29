package request

import (
	"github/jizi19911101/gin-vue-admin/server/model/autocode"
	"github/jizi19911101/gin-vue-admin/server/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}