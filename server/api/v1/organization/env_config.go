package organization

import (
	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/response"
	organization "github.com/jizi19911101/gin-vue-admin/server/model/organization"
	organizationReq "github.com/jizi19911101/gin-vue-admin/server/model/organization/request"
	organizationRes "github.com/jizi19911101/gin-vue-admin/server/model/organization/response"
	"github.com/jizi19911101/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

type EnvConfigApi struct {
}

var envConfigService = service.ServiceGroupApp.OrganizationServiceGroup.EnvConfigService

// CreateEnvConfig 创建EnvConfig
// @Tags EnvConfig
// @Summary 创建EnvConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EnvConfig true "创建EnvConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /envConfig/createEnvConfig [post]
func (envConfigApi *EnvConfigApi) CreateEnvConfig(c *gin.Context) {
	var envConfigReq organizationReq.EnvConfigReq
	_ = c.ShouldBindJSON(&envConfigReq)
	if err := global.Validate.Struct(envConfigReq); err != nil {
		global.GVA_LOG.Error("参数缺失", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}
	envConfig := envConfigApi.transferEnvconfig(envConfigReq)

	if err := envConfigService.CreateEnvConfig(envConfig); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEnvConfig 删除EnvConfig
// @Tags EnvConfig
// @Summary 删除EnvConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EnvConfig true "删除EnvConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /envConfig/deleteEnvConfig [delete]
func (envConfigApi *EnvConfigApi) DeleteEnvConfig(c *gin.Context) {
	var envConfigReq organizationReq.EnvConfigReq
	_ = c.ShouldBindJSON(&envConfigReq)
	envConfig := envConfigApi.transferEnvconfig(envConfigReq)
	if err := envConfigService.DeleteEnvConfig(envConfig); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEnvConfigByIds 批量删除EnvConfig
// @Tags EnvConfig
// @Summary 批量删除EnvConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EnvConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /envConfig/deleteEnvConfigByIds [delete]
func (envConfigApi *EnvConfigApi) DeleteEnvConfigByIds(c *gin.Context) {
	var IdsReq request.IdsReq
	_ = c.ShouldBindJSON(&IdsReq)
	if err := envConfigService.DeleteEnvConfigByIds(IdsReq); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEnvConfig 更新EnvConfig
// @Tags EnvConfig
// @Summary 更新EnvConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EnvConfig true "更新EnvConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /envConfig/updateEnvConfig [put]
func (envConfigApi *EnvConfigApi) UpdateEnvConfig(c *gin.Context) {
	var envConfigReq organizationReq.EnvConfigReq
	_ = c.ShouldBindJSON(&envConfigReq)
	if err := global.Validate.Struct(envConfigReq); err != nil {
		global.GVA_LOG.Error("参数缺失", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}
	envConfig := envConfigApi.transferEnvconfig(envConfigReq)
	if err := envConfigService.UpdateEnvConfig(envConfig); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEnvConfig 用id查询EnvConfig
// @Tags EnvConfig
// @Summary 用id查询EnvConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.EnvConfig true "用id查询EnvConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /envConfig/findEnvConfig [get]
func (envConfigApi *EnvConfigApi) FindEnvConfig(c *gin.Context) {
	var envConfigReq organizationReq.EnvConfigReq
	_ = c.ShouldBindQuery(&envConfigReq)
	if err, envConfig := envConfigService.GetEnvConfig(envConfigReq.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		envConfigRes := organizationRes.EnvConfigRes{
			ID:           envConfig.ID,
			Name:         envConfig.Name,
			Base_url:     envConfig.Base_url,
			Organization: envConfig.Organization,
		}
		response.OkWithData(gin.H{"envConfig": envConfigRes}, c)
	}
}

// GetEnvConfigList 分页获取EnvConfig列表
// @Tags EnvConfig
// @Summary 分页获取EnvConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.EnvConfigSearch true "分页获取EnvConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /envConfig/getEnvConfigList [get]
func (envConfigApi *EnvConfigApi) GetEnvConfigList(c *gin.Context) {
	var organizationReq organizationReq.EnvConfigSearch
	_ = c.ShouldBindQuery(&organizationReq)
	if err, list, total := envConfigService.GetEnvConfigInfoList(organizationReq); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		envConfigList := list.([]organization.EnvConfig)
		envConfigResList := make([]organizationRes.EnvConfigRes, 0)
		for _, envconfig := range envConfigList {
			envConfigResList = append(envConfigResList, organizationRes.EnvConfigRes{
				ID:           envconfig.ID,
				Name:         envconfig.Name,
				Base_url:     envconfig.Base_url,
				Organization: envconfig.Organization,
			})
		}

		response.OkWithDetailed(response.PageResult{
			List:     envConfigResList,
			Total:    total,
			Page:     organizationReq.Page,
			PageSize: organizationReq.PageSize,
		}, "获取成功", c)
	}
}

func (envConfigApi *EnvConfigApi) transferEnvconfig(envConfigReq organizationReq.EnvConfigReq) (evnConfig organization.EnvConfig) {
	//evnConfig = organization.EnvConfig{}
	evnConfig.ID = envConfigReq.ID
	evnConfig.Name = envConfigReq.Name
	evnConfig.Base_url = envConfigReq.Base_url
	evnConfig.Organization = envConfigReq.Organization
	return
}
