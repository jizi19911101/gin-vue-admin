package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EnvConfigRouter struct {
}

// InitEnvConfigRouter 初始化 EnvConfig 路由信息
func (s *EnvConfigRouter) InitEnvConfigRouter(Router *gin.RouterGroup) {
	envConfigRouter := Router.Group("envConfig").Use(middleware.OperationRecord())
	envConfigRouterWithoutRecord := Router.Group("envConfig")
	var envConfigApi = v1.ApiGroupApp.AutoCodeApiGroup.EnvConfigApi
	{
		envConfigRouter.POST("createEnvConfig", envConfigApi.CreateEnvConfig)   // 新建EnvConfig
		envConfigRouter.DELETE("deleteEnvConfig", envConfigApi.DeleteEnvConfig) // 删除EnvConfig
		envConfigRouter.DELETE("deleteEnvConfigByIds", envConfigApi.DeleteEnvConfigByIds) // 批量删除EnvConfig
		envConfigRouter.PUT("updateEnvConfig", envConfigApi.UpdateEnvConfig)    // 更新EnvConfig
	}
	{
		envConfigRouterWithoutRecord.GET("findEnvConfig", envConfigApi.FindEnvConfig)        // 根据ID获取EnvConfig
		envConfigRouterWithoutRecord.GET("getEnvConfigList", envConfigApi.GetEnvConfigList)  // 获取EnvConfig列表
	}
}
