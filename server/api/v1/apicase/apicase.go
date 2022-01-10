package apicase

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/apicase"
	apicaseReq "github.com/jizi19911101/gin-vue-admin/server/model/apicase/request"
	apicaseRes "github.com/jizi19911101/gin-vue-admin/server/model/apicase/response"
	"github.com/jizi19911101/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

type ApiCaseApi struct {
}

var apiCaseService = service.ServiceGroupApp.ApiCaseServiceGroup.ApiCaseService

// RunApiCase 跑接口测试用例
// @Tags RunApiCase
// @Summary 跑接口测试用例
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apicaseReq.RunApiCaseReq true "跑接口测试用例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"成功发起测试，稍后生成测试报告"}"
// @Router /apiCase/runApiCase [post]
func (apiCaseApi *ApiCaseApi) RunApiCase(c *gin.Context) {
	var runApiCaseReq apicaseReq.RunApiCaseReq
	_ = c.ShouldBindJSON(&runApiCaseReq)
	if err := global.Validate.Struct(&runApiCaseReq); err != nil {
		global.GVA_LOG.Error("参数缺失", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}
	if err := apiCaseService.RunApiCase(runApiCaseReq); err != nil {
		global.GVA_LOG.Error("发起测试出错", zap.Error(err))
		response.FailWithMessage("发起测试出错", c)
		return
	}
	response.OkWithMessage("成功发起测试，稍后生成测试报告", c)
}

func (apiCaseApi *ApiCaseApi) ModuleList(c *gin.Context) {
	var moduleListReq apicaseReq.ModuleSearch
	_ = c.ShouldBindQuery(&moduleListReq)
	if err, list, total := apiCaseService.ModuleList(moduleListReq); err != nil {
		global.GVA_LOG.Error("获取模块列表失败", zap.Error(err))
		response.FailWithMessage("获取模块列表失败", c)
	} else {
		moduleList := list.([]apicase.Module)
		moduleListRes := make([]apicaseRes.ModuleRes, 0)
		for i := range moduleList {
			moduleListRes = append(moduleListRes, apicaseRes.ModuleRes{
				ID:             moduleList[i].ID,
				Name:           moduleList[i].Name,
				OrganizationID: moduleList[i].OrganizationID,
			})
		}
		response.OkWithDetailed(response.PageResult{
			List:     moduleListRes,
			Total:    total,
			Page:     moduleListReq.Page,
			PageSize: moduleListReq.PageSize,
		}, "获取模块列表成功", c)
	}

}

func (apiCaseApi *ApiCaseApi) ApiList(c *gin.Context) {
	var apiListReq apicaseReq.ApiSearch
	_ = c.ShouldBindQuery(&apiListReq)
	if err, list, total := apiCaseService.ApiList(apiListReq); err != nil {
		global.GVA_LOG.Error("获取接口列表失败", zap.Error(err))
		response.FailWithMessage("获取接口列表失败", c)
	} else {
		apiList := list.([]apicase.Api)
		apiListRes := make([]apicaseRes.ApiRes, 0)
		for i := range apiList {
			apiListRes = append(apiListRes, apicaseRes.ApiRes{
				Name:           apiList[i].Name,
				Module:         apiList[i].Module,
				OrganizationID: apiList[i].OrganizationID,
			})
		}
		response.OkWithDetailed(response.PageResult{
			List:     apiListRes,
			Total:    total,
			Page:     apiListReq.Page,
			PageSize: apiListReq.PageSize,
		}, "获取接口列表成功", c)
	}

}

func (apiCaseApi *ApiCaseApi) CaseList(c *gin.Context) {
	response.OkWithData(gin.H{}, c)
}

func (apiCaseApi *ApiCaseApi) ReportList(c *gin.Context) {
	response.OkWithData(gin.H{}, c)
}
