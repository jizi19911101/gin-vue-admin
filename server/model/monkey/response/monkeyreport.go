package request

import "time"

type ReportRes struct {
	ID        uint
	Name      string    `json:"name" form:"name"  validate:"required"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt" `
}

//type HtmlRes struct {
//	Name      string    `json:"name" form:"name"  validate:"required"`
//	CreatedAt time.Time `json:"createdAt" form:"createdAt" `
//}
