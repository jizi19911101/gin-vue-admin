package organization

import (
	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/response"
	"github.com/jizi19911101/gin-vue-admin/server/model/organization"
	organizationReq "github.com/jizi19911101/gin-vue-admin/server/model/organization/request"
	organizationRes "github.com/jizi19911101/gin-vue-admin/server/model/organization/response"
	"github.com/jizi19911101/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

type OrganizationApi struct {
}

var organizationService = service.ServiceGroupApp.OrganizationServiceGroup.OrganizationService

// CreateOrganization 创建Organization
// @Tags Organization
// @Summary 创建Organization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body organizationReq.OrganizationReq true "创建Organization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organization/createOrganization [post]
func (organizationApi *OrganizationApi) CreateOrganization(c *gin.Context) {
	var organizationReq organizationReq.OrganizationReq
	_ = c.ShouldBindJSON(&organizationReq)
	if err := global.Validate.Struct(organizationReq); err != nil {
		global.GVA_LOG.Error("参数缺失", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}
	organization := organizationApi.transferOrganization(organizationReq)
	if err := organizationService.CreateOrganization(organization); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrganization 删除Organization
// @Tags Organization
// @Summary 删除Organization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body organizationReq.OrganizationReq true "删除Organization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /organization/deleteOrganization [delete]
func (organizationApi *OrganizationApi) DeleteOrganization(c *gin.Context) {
	var organizationReq organizationReq.OrganizationReq
	_ = c.ShouldBindJSON(&organizationReq)
	organization := organizationApi.transferOrganization(organizationReq)
	if err := organizationService.DeleteOrganization(organization); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrganizationByIds 批量删除
// @Tags Organization
// @Summary 批量删除Organization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Organization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /organization/deleteOrganizationByIds [delete]
func (organizationApi *OrganizationApi) DeleteOrganizationByIds(c *gin.Context) {
	var IdsReq request.IdsReq
	_ = c.ShouldBindJSON(&IdsReq)
	if err := organizationService.DeleteOrganizationByIds(IdsReq); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrganization 更新Organization
// @Tags Organization
// @Summary 更新Organization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body organizationReq.OrganizationReq true "更新Organization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /organization/updateOrganization [put]
func (organizationApi *OrganizationApi) UpdateOrganization(c *gin.Context) {
	var organizationReq organizationReq.OrganizationReq
	_ = c.ShouldBindJSON(&organizationReq)
	if err := global.Validate.Struct(organizationReq); err != nil {
		global.GVA_LOG.Error("参数缺失", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}
	organization := organizationApi.transferOrganization(organizationReq)
	if err := organizationService.UpdateOrganization(organization); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrganization 用id查询Organization
// @Tags Organization
// @Summary 用id查询Organization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query organizationReq.OrganizationReq true "用id查询Organization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /organization/findOrganization [get]
func (organizationApi *OrganizationApi) FindOrganization(c *gin.Context) {
	var organizationReq organizationReq.OrganizationReq
	_ = c.ShouldBindQuery(&organizationReq)
	if err, organization := organizationService.GetOrganization(organizationReq.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		organizationRes := organizationRes.OrganizationRes{
			ID:   organization.ID,
			Name: organization.Name,
		}
		response.OkWithData(gin.H{"organization": organizationRes}, c)
	}

}

// GetOrganizationList 分页获取Organization列表
// @Tags Organization
// @Summary 分页获取Organization列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query organizationReq.OrganizationSearch true "分页获取Organization列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organization/getOrganizationList [get]
func (organizationApi *OrganizationApi) GetOrganizationList(c *gin.Context) {
	var organizationReq organizationReq.OrganizationSearch
	_ = c.ShouldBindQuery(&organizationReq)
	if err, list, total := organizationService.GetOrganizationInfoList(organizationReq); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		organizationList := list.([]organization.Organization)
		organizationResList := make([]organizationRes.OrganizationRes, 0)
		for _, organization := range organizationList {
			organizationResList = append(organizationResList, organizationRes.OrganizationRes{
				ID:   organization.ID,
				Name: organization.Name,
			})
		}
		response.OkWithDetailed(response.PageResult{
			List:     organizationResList,
			Total:    total,
			Page:     organizationReq.Page,
			PageSize: organizationReq.PageSize,
		}, "获取成功", c)
	}
}

func (organizationApi *OrganizationApi) transferOrganization(organizationReq organizationReq.OrganizationReq) (organization organization.Organization) {
	organization.Name = organizationReq.Name
	organization.ID = organizationReq.ID

	return
}
