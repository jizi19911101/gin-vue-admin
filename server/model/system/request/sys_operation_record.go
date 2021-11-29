package request

import (
	"github/jizi19911101/gin-vue-admin/server/model/common/request"
	"github/jizi19911101/gin-vue-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
