package project

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jizi19911101/gin-vue-admin/server/api/v1"
	"github.com/jizi19911101/gin-vue-admin/server/middleware"
)

type ProjectRouter struct {
}

// InitProjectRouter 初始化 Project 路由信息
func (s *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup) {
	projectRouter := Router.Group("organization").Use(middleware.OperationRecord())
	projectRouterWithoutRecord := Router.Group("organization")
	var organizationApi = v1.ApiGroupApp.OrganizationGroup.OrganizationApi
	{
		projectRouter.POST("createProject", organizationApi.CreateProject)             // 新建Project
		projectRouter.DELETE("deleteProject", organizationApi.DeleteProject)           // 删除Project
		projectRouter.DELETE("deleteProjectByIds", organizationApi.DeleteProjectByIds) // 批量删除Project
		projectRouter.PUT("updateProject", organizationApi.UpdateProject)              // 更新Project
	}
	{
		projectRouterWithoutRecord.GET("findProject", organizationApi.FindProject)       // 根据ID获取Project
		projectRouterWithoutRecord.GET("getProjectList", organizationApi.GetProjectList) // 获取Project列表
	}
}
