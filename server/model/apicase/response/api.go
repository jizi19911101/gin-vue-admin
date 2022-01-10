package response

type ApiRes struct {
	Name           string `json:"name" form:"name"`
	Module         string `json:"module" form:"module"`
	OrganizationID uint   `json:"organizationID" form:"organizationID"`
}
