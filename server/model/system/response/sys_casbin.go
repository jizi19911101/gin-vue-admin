package response

import (
	"github.com/jizi19911101/gin-vue-admin/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
