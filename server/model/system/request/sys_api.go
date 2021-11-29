package request

import (
	"github/jizi19911101/gin-vue-admin/server/model/common/request"
	"github/jizi19911101/gin-vue-admin/server/model/system"
)

// api分页条件查询及排序结构体
type SearchApiParams struct {
	system.SysApi
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
