package request

import (
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
