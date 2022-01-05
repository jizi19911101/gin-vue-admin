package sync

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jizi19911101/gin-vue-admin/server/api/v1"
)

type SyncRouter struct {
}

func (s *SyncRouter) InitSyncRouter(Router *gin.RouterGroup) {
	syncRouter := Router.Group("sync")
	//apiTestcaseRouterWithRecord := Router.Group("apiTestcase")
	var syncApi = v1.ApiGroupApp.SyncGroup
	{
		syncRouter.GET("syncApiTestcase", syncApi.SyncApiTestCaseApi)

	}
}
