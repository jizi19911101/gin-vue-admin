package request

type ReportReq struct {
	Name string `json:"name" form:"name"  validate:"required"`
	Url  string `json:"url" form:"url"  validate:"required"`
}
