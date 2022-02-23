package response

import "time"

type ReportRes struct {
	ID             uint
	Name           string    `json:"name" form:"name"  `
	Url            string    `json:"url" form:"url" `
	Env            string    `json:"env" form:"env"  `
	Description    string    `json:"description" form:"description"  `
	OrganizationID uint      `json:"organizationID" form:"organizationID" `
	CreatedAt      time.Time `json:"createdAt" form:"createdAt" `
}
