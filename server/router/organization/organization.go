package organization

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jizi19911101/gin-vue-admin/server/api/v1"
	"github.com/jizi19911101/gin-vue-admin/server/middleware"
)

type OrganizationRouter struct {
}

// InitProjectRouter 初始化 Project 路由信息
func (s *OrganizationRouter) InitProjectRouter(Router *gin.RouterGroup) {
	organizationRouter := Router.Group("organization").Use(middleware.OperationRecord())
	organizationRouterWithoutRecord := Router.Group("organization")
	var organizationApi = v1.ApiGroupApp.OrganizationGroup.OrganizationApi
	{
		organizationRouter.POST("createOrganization", organizationApi.CreateOrganization)             // 新建Project
		organizationRouter.DELETE("deleteOrganization", organizationApi.DeleteOrganization)           // 删除Project
		organizationRouter.DELETE("deleteOrganizationByIds", organizationApi.DeleteOrganizationByIds) // 批量删除Project
		organizationRouter.PUT("updateOrganization", organizationApi.UpdateOrganization)              // 更新Project
	}
	{
		organizationRouterWithoutRecord.GET("findOrganization", organizationApi.FindOrganization)       // 根据ID获取Project
		organizationRouterWithoutRecord.GET("getOrganizationList", organizationApi.GetOrganizationList) // 获取Project列表
	}
}
