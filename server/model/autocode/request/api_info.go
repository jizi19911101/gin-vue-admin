package request

import (
	"github.com/jizi19911101/gin-vue-admin/server/model/api_test"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
)

type ApiInfoSearch struct {
	api_test.ApiInfo
	request.PageInfo
}
