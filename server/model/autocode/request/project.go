package request

import (
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/project"
)

type ProjectSearch struct {
	project.Project
	request.PageInfo
}