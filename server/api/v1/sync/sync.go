package sync

import (
	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/apicase"
	apicaseReq "github.com/jizi19911101/gin-vue-admin/server/model/apicase/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/response"
	"github.com/jizi19911101/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

type SyncApi struct {
}

var syncService = service.ServiceGroupApp.SyncServiceGroup.SyncService

// SyncApiTestCaseApi 同步并解析接口自动化代码
// @Tags SyncApiTestCaseApi
// @Summary 同步并解析接口自动化代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"同步并解析接口自动化代码成功！"}"
// @Router /sync/syncApiTestcase [get]
func (syncApi *SyncApi) SyncApiTestCaseApi(c *gin.Context) {
	if err := syncService.SyncApiTestCase(); err != nil {
		global.GVA_LOG.Error("同步并解析接口自动化代码！", zap.Error(err))
		response.FailWithMessage("同步并解析接口自动化代码失败！", c)
	} else {
		response.OkWithMessage("同步并解析接口自动化代码成功！", c)
	}
}

// SyncApiTestReportApi 同步接口测试报告
// @Tags SyncApiTestReportApi
// @Summary 同步接口测试报告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"同步接口测试报告成功！"}"
// @Router /sync/syncApiTestReport [get]
func (syncApi *SyncApi) SyncApiTestReportApi(c *gin.Context) {
	var reportReq apicaseReq.ReportReq
	_ = c.ShouldBindQuery(&reportReq)
	if err := global.Validate.Struct(&reportReq); err != nil {
		global.GVA_LOG.Error("参数缺失", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}

	report := apicase.Report{
		Name: reportReq.Name,
		Url:  reportReq.Url,
	}
	if err := syncService.SyncApiTestReport(report); err != nil {
		global.GVA_LOG.Error("同步接口测试报告失败！", zap.Error(err))
		response.FailWithMessage("同步接口测试报告失败！", c)
	} else {
		response.OkWithMessage("同步接口测试报告成功！", c)
	}
}
