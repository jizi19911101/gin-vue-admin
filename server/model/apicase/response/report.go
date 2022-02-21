package response

import "time"

type ReportRes struct {
	ID             uint
	Name           string    `json:"name" form:"name"  `
	Url            string    `json:"url" form:"url" `
	OrganizationID uint      `json:"organizationID" form:"organizationID" `
	CreatedAt      time.Time `json:"createdAt" form:"createdAt" `
}
