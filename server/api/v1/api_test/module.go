package api_test

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/jizi19911101/gin-vue-admin/server/model/autocode/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/response"
	"github.com/jizi19911101/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

type ModuleApi struct {
}

var moduleService = service.ServiceGroupApp.AutoCodeServiceGroup.ModuleService

// CreateModule 创建Module
// @Tags Module
// @Summary 创建Module
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Module true "创建Module"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /module/createModule [post]
func (moduleApi *ModuleApi) CreateModule(c *gin.Context) {
	var module autocode.Module
	_ = c.ShouldBindJSON(&module)
	if err := moduleService.CreateModule(module); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteModule 删除Module
// @Tags Module
// @Summary 删除Module
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Module true "删除Module"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /module/deleteModule [delete]
func (moduleApi *ModuleApi) DeleteModule(c *gin.Context) {
	var module autocode.Module
	_ = c.ShouldBindJSON(&module)
	if err := moduleService.DeleteModule(module); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteModuleByIds 批量删除Module
// @Tags Module
// @Summary 批量删除Module
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Module"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /module/deleteModuleByIds [delete]
func (moduleApi *ModuleApi) DeleteModuleByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := moduleService.DeleteModuleByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateModule 更新Module
// @Tags Module
// @Summary 更新Module
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Module true "更新Module"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /module/updateModule [put]
func (moduleApi *ModuleApi) UpdateModule(c *gin.Context) {
	var module autocode.Module
	_ = c.ShouldBindJSON(&module)
	fmt.Println("12122module", module)
	if err := moduleService.UpdateModule(module); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindModule 用id查询Module
// @Tags Module
// @Summary 用id查询Module
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Module true "用id查询Module"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /module/findModule [get]
func (moduleApi *ModuleApi) FindModule(c *gin.Context) {
	var module autocode.Module
	_ = c.ShouldBindQuery(&module)
	if err, remodule := moduleService.GetModule(module.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remodule": remodule}, c)
	}
}

// GetModuleList 分页获取Module列表
// @Tags Module
// @Summary 分页获取Module列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ModuleSearch true "分页获取Module列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /module/getModuleList [get]
func (moduleApi *ModuleApi) GetModuleList(c *gin.Context) {
	var pageInfo autocodeReq.ModuleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := moduleService.GetModuleInfoList(pageInfo); err != nil {
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
