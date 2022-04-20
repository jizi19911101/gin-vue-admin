package response

type ApiCaseRes struct {
	ID             uint
	OrganizationID uint   `json:"organizationID" form:"organizationID" `
	Module         string `json:"module" form:"module"  `
	Api            string `json:"api" form:"api"  `
	Name           string `json:"name" form:"name"  `
	Title          string `json:"title" form:"title"  `
}
