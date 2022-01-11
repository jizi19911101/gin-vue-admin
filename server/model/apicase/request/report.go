package request

import "github.com/jizi19911101/gin-vue-admin/server/model/common/request"

type ReportReq struct {
	Name string `json:"name" form:"name"  validate:"required"`
	Url  string `json:"url" form:"url"  validate:"required"`
}

type ReportSearch struct {
	request.PageInfo
	Name string `json:"name" form:"name"  `
}
