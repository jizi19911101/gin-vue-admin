package response

//type OrganizationSearch struct {
//	organization.Organization
//	request.PageInfo
//}

type OrganizationRes struct {
	ID   uint
	Name string `json:"name" form:"name"  validate:"required"`
}
