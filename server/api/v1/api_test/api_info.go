package api_test

import (
	"strings"

	apiInfoReq "github.com/jizi19911101/gin-vue-admin/server/model/api_test/request"
	apiInfoRes "github.com/jizi19911101/gin-vue-admin/server/model/api_test/response"

	"github.com/jizi19911101/gin-vue-admin/server/model/api_test"

	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	autocodeReq "github.com/jizi19911101/gin-vue-admin/server/model/autocode/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/response"
	"github.com/jizi19911101/gin-vue-admin/server/service"
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
	var apiInfoRequest apiInfoReq.ApiInfoRequest
	_ = c.ShouldBindJSON(&apiInfoRequest)
	//global.Validate = validator.New()
	if err := global.Validate.Struct(&apiInfoRequest); err != nil {
		global.GVA_LOG.Error("参数缺失!", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}
	apiInfo := apiInfoApi.transferRequest(apiInfoRequest)
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
	var apiInfo api_test.ApiInfo
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
	//var apiInfo autocode.ApiInfo
	var apiInfoRequest apiInfoReq.ApiInfoRequest
	_ = c.ShouldBindJSON(&apiInfoRequest)
	if err := global.Validate.Struct(&apiInfoRequest); err != nil {
		global.GVA_LOG.Error("参数缺失!", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}
	apiInfo := apiInfoApi.transferRequest(apiInfoRequest)
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
	var apiInfo api_test.ApiInfo
	_ = c.ShouldBindQuery(&apiInfo)
	if err, reapiInfo := apiInfoService.GetApiInfo(apiInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		resApiInfo := apiInfoApi.transferResponse(reapiInfo)
		response.OkWithData(gin.H{"reapiInfo": resApiInfo}, c)
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
		apiList := list.([]api_test.ApiInfo)
		resApiInfoList := make([]apiInfoRes.ApiInfoReSponse, 0, len(apiList))

		for _, api := range apiList {
			resApiInfo := apiInfoApi.transferResponse(api)
			resApiInfoList = append(resApiInfoList, resApiInfo)
		}
		response.OkWithDetailed(response.PageResult{
			List:     resApiInfoList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (apiInfoApi *ApiInfoApi) transferRequest(apiInfoRequest apiInfoReq.ApiInfoRequest) api_test.ApiInfo {
	apiInfo := api_test.ApiInfo{
		GVA_MODEL: global.GVA_MODEL{
			ID: apiInfoRequest.ID,
		},
		Url:     apiInfoRequest.Url,
		Name:    apiInfoRequest.Name,
		Method:  apiInfoRequest.Method,
		Module:  apiInfoRequest.Module,
		Project: apiInfoRequest.Project,
	}
	apiInfo.Params = strings.Join(apiInfoRequest.Params, ",")
	return apiInfo
}

func (apiInfoApi *ApiInfoApi) transferResponse(apiInfo api_test.ApiInfo) apiInfoRes.ApiInfoReSponse {
	apiInfoResponse := apiInfoRes.ApiInfoReSponse{
		ID:      apiInfo.ID,
		Url:     apiInfo.Url,
		Name:    apiInfo.Name,
		Method:  apiInfo.Method,
		Module:  apiInfo.Module,
		Project: apiInfo.Project,
	}

	apiInfoResponse.Params = strings.Split(apiInfo.Params, ",")
	return apiInfoResponse
}
