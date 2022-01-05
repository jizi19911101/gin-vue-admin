package sync

import (
	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/response"
	"github.com/jizi19911101/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

type SyncApi struct {
}

var syncService = service.ServiceGroupApp.SyncGroup.SyncService

// ApiTestcaseCode 同步并解析接口自动化代码
// @Tags ApiTestcaseCode
// @Summary 同步并解析接口自动化代码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"同步并解析接口自动化代码成功！"}"
// @Router /apiTestcase/apiTestcaseCode [get]
func (syncApi *SyncApi) SyncApiTestCaseApi(c *gin.Context) {
	if err := syncService.SyncApiTestCase(); err != nil {
		global.GVA_LOG.Error("同步并解析接口自动化代码！", zap.Error(err))
		response.FailWithMessage("同步并解析接口自动化代码失败！", c)
	} else {
		response.OkWithMessage("同步并解析接口自动化代码成功！", c)
	}
}
