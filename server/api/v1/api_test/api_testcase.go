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

// CreateApiInfo 创建ApiInfo
// @Tags ApiInfo
// @Summary 创建ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ApiInfo true "创建ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiInfo/createApiInfo [post]
func (apiTestcaseApi *ApiTestcaseApi) ApiTestcaseCodeApi(c *gin.Context) {
	if err := apiTestcaseService.ApiTestcaseCode(); err != nil {
		global.GVA_LOG.Error("拉取代码失败！", zap.Error(err))
		response.FailWithMessage("拉取代码失败！", c)
	} else {
		response.OkWithMessage("拉取代码成功！", c)
	}
}
