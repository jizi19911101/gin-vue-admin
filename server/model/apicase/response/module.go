package response

type ModuleRes struct {
	ID             uint
	Name           string `json:"name" form:"name" `
	OrganizationID uint   `json:"organizationID" form:"organizationID" `
}
