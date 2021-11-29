package response

import "github.com/jizi19911101/gin-vue-admin/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
