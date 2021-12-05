package request

import (
	"github.com/jizi19911101/gin-vue-admin/server/model/api_test"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
)

type ModuleSearch struct {
	api_test.Module
	request.PageInfo
}
