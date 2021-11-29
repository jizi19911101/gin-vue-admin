package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiInfoRouter struct {
}

// InitApiInfoRouter 初始化 ApiInfo 路由信息
func (s *ApiInfoRouter) InitApiInfoRouter(Router *gin.RouterGroup) {
	apiInfoRouter := Router.Group("apiInfo").Use(middleware.OperationRecord())
	apiInfoRouterWithoutRecord := Router.Group("apiInfo")
	var apiInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.ApiInfoApi
	{
		apiInfoRouter.POST("createApiInfo", apiInfoApi.CreateApiInfo)   // 新建ApiInfo
		apiInfoRouter.DELETE("deleteApiInfo", apiInfoApi.DeleteApiInfo) // 删除ApiInfo
		apiInfoRouter.DELETE("deleteApiInfoByIds", apiInfoApi.DeleteApiInfoByIds) // 批量删除ApiInfo
		apiInfoRouter.PUT("updateApiInfo", apiInfoApi.UpdateApiInfo)    // 更新ApiInfo
	}
	{
		apiInfoRouterWithoutRecord.GET("findApiInfo", apiInfoApi.FindApiInfo)        // 根据ID获取ApiInfo
		apiInfoRouterWithoutRecord.GET("getApiInfoList", apiInfoApi.GetApiInfoList)  // 获取ApiInfo列表
	}
}
