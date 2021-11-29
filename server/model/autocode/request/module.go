package request

import (
	"github/jizi19911101/gin-vue-admin/server/model/autocode"
	"github/jizi19911101/gin-vue-admin/server/model/common/request"
)

type ModuleSearch struct{
    autocode.Module
    request.PageInfo
}