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
		apiCaseRouter.GET("moduleList", apiCaseApi.ModuleList)
		apiCaseRouter.GET("apiList", apiCaseApi.ApiList)
		apiCaseRouter.GET("caseList", apiCaseApi.CaseList)
		apiCaseRouter.GET("reportList", apiCaseApi.ReportList)
		//apiCaseRouter.GET("searchReport", apiCaseApi.SearchReport)
	}
}
