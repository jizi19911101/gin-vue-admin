package apicase

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	apicaseReq "github.com/jizi19911101/gin-vue-admin/server/model/apicase/request"
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
	return
}

func (apiCaseApi *ApiCaseApi) ModuleInfo(c *gin.Context) {
	response.OkWithData(gin.H{}, c)
	return
}

func (apiCaseApi *ApiCaseApi) ApiInfo(c *gin.Context) {
	response.OkWithData(gin.H{}, c)
	return
}

func (apiCaseApi *ApiCaseApi) CaseInfo(c *gin.Context) {
	response.OkWithData(gin.H{}, c)
	return
}

func (apiCaseApi *ApiCaseApi) Report(c *gin.Context) {
	response.OkWithData(gin.H{}, c)
	return
}

func (apiCaseApi *ApiCaseApi) SearchReport(c *gin.Context) {
	response.OkWithData(gin.H{}, c)
	return
}
