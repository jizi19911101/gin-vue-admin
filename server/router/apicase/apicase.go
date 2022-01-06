package apicase

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jizi19911101/gin-vue-admin/server/api/v1"
	"github.com/jizi19911101/gin-vue-admin/server/middleware"
)

type ApiCaseRouter struct {
}

func (s *ApiCaseRouter) InitApiCaseRouter(Router *gin.RouterGroup) {
	apiCaseRouter := Router.Group("apiCase").Use(middleware.OperationRecord())
	var apiCaseApi = v1.ApiGroupApp.ApiCaseGroup.ApiCaseApi
	{
		apiCaseRouter.POST("runApiCase", apiCaseApi.RunApiCase)
		apiCaseRouter.GET("moduleInfo", apiCaseApi.ModuleInfo)
		apiCaseRouter.GET("apiInfo", apiCaseApi.ApiInfo)
		apiCaseRouter.GET("caseInfo", apiCaseApi.CaseInfo)
		apiCaseRouter.GET("report", apiCaseApi.Report)
		apiCaseRouter.GET("searchReport", apiCaseApi.SearchReport)
	}
}
