package monkey

import (
	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/response"
	"github.com/jizi19911101/gin-vue-admin/server/model/monkey"
	monkeyReq "github.com/jizi19911101/gin-vue-admin/server/model/monkey/request"
	monkeyRes "github.com/jizi19911101/gin-vue-admin/server/model/monkey/response"
	"github.com/jizi19911101/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

type MonkeyApi struct {
}

var monkeyService = service.ServiceGroupApp.MonkeyServiceGroup.MonkeyService

func (monkeyApi *MonkeyApi) StartMonkeyApi(c *gin.Context) {
	var startMonkeyReq monkeyReq.StartMonkeyReq
	_ = c.ShouldBindJSON(&startMonkeyReq)
	if err := global.Validate.Struct(&startMonkeyReq); err != nil {
		global.GVA_LOG.Error("参数缺失", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}
	err := monkeyService.StartMonkey(startMonkeyReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return

	}

	response.OkWithMessage("成功发起monkey测试，稍后生成测试报告", c)

}

// ReportList 获取报告列表
// @Tags ReportList
// @Summary 获取报告列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body apicaseReq.ReportSearch true "获取报告列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取报告列表成功!"}"
// @Router /apiCase/reportList [get]
func (monkeyApi *MonkeyApi) ReportList(c *gin.Context) {
	var reportListReq monkeyReq.ReportSearch
	_ = c.ShouldBindQuery(&reportListReq)
	if err, list, total := monkeyService.ReportList(reportListReq); err != nil {
		global.GVA_LOG.Error("获取报告列表失败", zap.Error(err))
		response.FailWithMessage("获取报告列表失败", c)
	} else {
		reportList := list.([]monkey.MonkeyReport)
		reportListRes := make([]monkeyRes.ReportRes, 0)
		for i := range reportList {
			reportListRes = append(reportListRes, monkeyRes.ReportRes{
				ID:        reportList[i].ID,
				Name:      reportList[i].Name,
				CreatedAt: reportList[i].CreatedAt,
			})
		}
		response.OkWithDetailed(response.PageResult{
			List:     reportListRes,
			Total:    total,
			Page:     reportListReq.Page,
			PageSize: reportListReq.PageSize,
		}, "获取报告列表成功!", c)
	}

}
