package request

import "github.com/jizi19911101/gin-vue-admin/server/model/common/request"

type ReportSearch struct {
	request.PageInfo
	Name string `json:"name" form:"name"  `
}

type HtmlReq struct {
	ID uint
}
