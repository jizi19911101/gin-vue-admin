package organization

import (
	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/response"
	"github.com/jizi19911101/gin-vue-admin/server/model/organization"
	organizationReq "github.com/jizi19911101/gin-vue-admin/server/model/organization/request"
	"github.com/jizi19911101/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

type OrganizationApi struct {
}

var projectService = service.ServiceGroupApp.ProjectServiceGroup.ProjectService

// CreateProject 创建Project
// @Tags Project
// @Summary 创建Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Project true "创建Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organization/createProject [post]
func (organizationApi *OrganizationApi) CreateProject(c *gin.Context) {
	var organizationReq organizationReq.OrganizationReq
	_ = c.ShouldBindJSON(&organizationReq)
	if err := global.Validate.Struct(organizationReq); err != nil {
		global.GVA_LOG.Error("参数缺失", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}
	organization := organizationApi.transferOrganization(organizationReq)
	if err := projectService.CreateProject(organization); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProject 删除Project
// @Tags Project
// @Summary 删除Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Project true "删除Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /organization/deleteProject [delete]
func (organizationApi *OrganizationApi) DeleteProject(c *gin.Context) {
	var organizationReq organizationReq.OrganizationReq
	_ = c.ShouldBindJSON(&organizationReq)
	organization := organizationApi.transferOrganization(organizationReq)
	if err := projectService.DeleteProject(organization); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProjectByIds 批量删除Project
// @Tags Project
// @Summary 批量删除Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /organization/deleteProjectByIds [delete]
func (organizationApi *OrganizationApi) DeleteProjectByIds(c *gin.Context) {
	var IdsReq request.IdsReq
	_ = c.ShouldBindJSON(&IdsReq)
	if err := projectService.DeleteProjectByIds(IdsReq); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProject 更新Project
// @Tags Project
// @Summary 更新Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Project true "更新Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /organization/updateProject [put]
func (organizationApi *OrganizationApi) UpdateProject(c *gin.Context) {
	var organizationReq organizationReq.OrganizationReq
	_ = c.ShouldBindJSON(&organizationReq)
	if err := global.Validate.Struct(organizationReq); err != nil {
		global.GVA_LOG.Error("参数缺失", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}
	organization := organizationApi.transferOrganization(organizationReq)
	if err := projectService.UpdateProject(organization); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProject 用id查询Project
// @Tags Project
// @Summary 用id查询Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Project true "用id查询Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /organization/findProject [get]
func (organizationApi *OrganizationApi) FindProject(c *gin.Context) {
	var organizationReq organizationReq.OrganizationReq
	_ = c.ShouldBindQuery(&organizationReq)
	if err, project := projectService.GetProject(organizationReq.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"organization": project}, c)
	}
}

// GetProjectList 分页获取Project列表
// @Tags Project
// @Summary 分页获取Project列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ProjectSearch true "分页获取Project列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organization/getProjectList [get]
func (organizationApi *OrganizationApi) GetProjectList(c *gin.Context) {
	var pageInfo organizationReq.OrganizationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := projectService.GetProjectInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (organizationApi *OrganizationApi) transferOrganization(organizationReq organizationReq.OrganizationReq) (organization organization.Organization) {
	organization.Name = organizationReq.Name
	return
}
