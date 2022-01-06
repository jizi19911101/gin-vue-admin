package request

type RunApiCaseReq struct {
	Organization string `json:"organization" form:"organization"  validate:"required"`
	Module       string `json:"module" form:"module"  `
	Api          string `json:"api" form:"api"  `
	Case         string `json:"case" form:"case"  `
	Env          string `json:"env" form:"env"  validate:"required"`
}
