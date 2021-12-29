package api_test

import (
	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/response"
	"github.com/jizi19911101/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

type ApiTestcaseApi struct {
}

var apiTestcaseService = service.ServiceGroupApp.AutoCodeServiceGroup.ApiTestcaseService

// ApiTestcaseCode 拉取接口自动化代码
// @Tags ApiTestcaseCode
// @Summary 拉取接口自动化代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body  true "拉取接口自动化代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉取代码成功！"}"
// @Router /apiTestcase/apiTestcaseCode [get]
func (apiTestcaseApi *ApiTestcaseApi) ApiTestcaseCodeApi(c *gin.Context) {
	if err := apiTestcaseService.ApiTestcaseCode(); err != nil {
		global.GVA_LOG.Error("拉取接口自动化代码失败！", zap.Error(err))
		response.FailWithMessage("拉取接口自动化代码失败！", c)
	} else {
		response.OkWithMessage("拉取接口自动化代码成功！", c)
	}
}

func (apiTestcaseApi *ApiTestcaseApi) ParseApiTestcaseApi(c *gin.Context) {
	if err := apiTestcaseService.ParseApiTestcaseApi(); err != nil {
		global.GVA_LOG.Error("解析接口自动化代码失败！", zap.Error(err))
		response.FailWithMessage("解析接口自动化代码失败！", c)
	} else {
		response.OkWithMessage("解析接口自动化代码成功！", c)
	}
}