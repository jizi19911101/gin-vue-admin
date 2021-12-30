package api_test

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jizi19911101/gin-vue-admin/server/api/v1"
	"github.com/jizi19911101/gin-vue-admin/server/middleware"
)

type ApiTestcaseRouter struct {
}

func (s *ApiTestcaseRouter) InitApiTestcaseRouter(Router *gin.RouterGroup) {
	apiTestcaseRouter := Router.Group("apiTestcase").Use(middleware.OperationRecord())
	//apiTestcaseRouterWithRecord := Router.Group("apiTestcase")
	var apiTestcaseApi = v1.ApiGroupApp.ApiTestGroup
	{
		apiTestcaseRouter.GET("apiTestcaseCode", apiTestcaseApi.ApiTestcaseCodeApi)
		apiTestcaseRouter.GET("parseApiTestcaseModule", apiTestcaseApi.ParseApiTestcaseModuleApi)
		apiTestcaseRouter.GET("parseApiTestcaseApi", apiTestcaseApi.ParseApiTestcaseApiApi)
		apiTestcaseRouter.GET("parseApiTestcase", apiTestcaseApi.ParseApiTestcaseApi)
	}
}
