package response

//type EnvConfigSearch struct {
//	organization.EnvConfig
//	request.PageInfo
//}

type EnvConfigRes struct {
	ID           uint
	Name         string `json:"name" form:"name"  validate:"required"`
	Base_url     string `json:"base_url" form:"base_url"   validate:"required"`
	Organization string `json:"organization" form:"organization"   validate:"required"`
}
