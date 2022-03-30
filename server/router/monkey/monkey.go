package monkey

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jizi19911101/gin-vue-admin/server/api/v1"
	"github.com/jizi19911101/gin-vue-admin/server/middleware"
)

type MonkeyRouter struct {
}

func (s *MonkeyRouter) InitMonkeyRouter(Router *gin.RouterGroup) {
	monkeyRouter := Router.Group("monkey").Use(middleware.OperationRecord())
	var monkeyApi = v1.ApiGroupApp.MonkeyGroup
	{
		monkeyRouter.POST("startMonkey", monkeyApi.StartMonkeyApi)
		monkeyRouter.GET("reportList", monkeyApi.ReportList)
		monkeyRouter.GET("reportContent", monkeyApi.ReportContent)

	}
}
