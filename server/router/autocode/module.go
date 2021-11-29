package autocode

import (
	"github.com/jizi19911101/gin-vue-admin/server/api/v1"
	"github.com/jizi19911101/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ModuleRouter struct {
}

// InitModuleRouter 初始化 Module 路由信息
func (s *ModuleRouter) InitModuleRouter(Router *gin.RouterGroup) {
	moduleRouter := Router.Group("module").Use(middleware.OperationRecord())
	moduleRouterWithoutRecord := Router.Group("module")
	var moduleApi = v1.ApiGroupApp.AutoCodeApiGroup.ModuleApi
	{
		moduleRouter.POST("createModule", moduleApi.CreateModule)   // 新建Module
		moduleRouter.DELETE("deleteModule", moduleApi.DeleteModule) // 删除Module
		moduleRouter.DELETE("deleteModuleByIds", moduleApi.DeleteModuleByIds) // 批量删除Module
		moduleRouter.PUT("updateModule", moduleApi.UpdateModule)    // 更新Module
	}
	{
		moduleRouterWithoutRecord.GET("findModule", moduleApi.FindModule)        // 根据ID获取Module
		moduleRouterWithoutRecord.GET("getModuleList", moduleApi.GetModuleList)  // 获取Module列表
	}
}
