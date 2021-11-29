package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ApiInfoApi struct {
}

var apiInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.ApiInfoService


// CreateApiInfo 创建ApiInfo
// @Tags ApiInfo
// @Summary 创建ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ApiInfo true "创建ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiInfo/createApiInfo [post]
func (apiInfoApi *ApiInfoApi) CreateApiInfo(c *gin.Context) {
	var apiInfo autocode.ApiInfo
	_ = c.ShouldBindJSON(&apiInfo)
	if err := apiInfoService.CreateApiInfo(apiInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApiInfo 删除ApiInfo
// @Tags ApiInfo
// @Summary 删除ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ApiInfo true "删除ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiInfo/deleteApiInfo [delete]
func (apiInfoApi *ApiInfoApi) DeleteApiInfo(c *gin.Context) {
	var apiInfo autocode.ApiInfo
	_ = c.ShouldBindJSON(&apiInfo)
	if err := apiInfoService.DeleteApiInfo(apiInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteApiInfoByIds 批量删除ApiInfo
// @Tags ApiInfo
// @Summary 批量删除ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /apiInfo/deleteApiInfoByIds [delete]
func (apiInfoApi *ApiInfoApi) DeleteApiInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := apiInfoService.DeleteApiInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateApiInfo 更新ApiInfo
// @Tags ApiInfo
// @Summary 更新ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ApiInfo true "更新ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apiInfo/updateApiInfo [put]
func (apiInfoApi *ApiInfoApi) UpdateApiInfo(c *gin.Context) {
	var apiInfo autocode.ApiInfo
	_ = c.ShouldBindJSON(&apiInfo)
	if err := apiInfoService.UpdateApiInfo(apiInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindApiInfo 用id查询ApiInfo
// @Tags ApiInfo
// @Summary 用id查询ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ApiInfo true "用id查询ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apiInfo/findApiInfo [get]
func (apiInfoApi *ApiInfoApi) FindApiInfo(c *gin.Context) {
	var apiInfo autocode.ApiInfo
	_ = c.ShouldBindQuery(&apiInfo)
	if err, reapiInfo := apiInfoService.GetApiInfo(apiInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapiInfo": reapiInfo}, c)
	}
}

// GetApiInfoList 分页获取ApiInfo列表
// @Tags ApiInfo
// @Summary 分页获取ApiInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ApiInfoSearch true "分页获取ApiInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiInfo/getApiInfoList [get]
func (apiInfoApi *ApiInfoApi) GetApiInfoList(c *gin.Context) {
	var pageInfo autocodeReq.ApiInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := apiInfoService.GetApiInfoInfoList(pageInfo); err != nil {
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
