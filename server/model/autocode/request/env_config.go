package request

import (
	"github/jizi19911101/gin-vue-admin/server/model/autocode"
	"github/jizi19911101/gin-vue-admin/server/model/common/request"
)

type EnvConfigSearch struct{
    autocode.EnvConfig
    request.PageInfo
}