package response

import "github.com/jizi19911101/gin-vue-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
