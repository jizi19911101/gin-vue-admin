package monkey

import (
	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/response"
	monkeyReq "github.com/jizi19911101/gin-vue-admin/server/model/monkey/request"
	"go.uber.org/zap"
)

type MonkeyApi struct {
}

func (monkeyApi *MonkeyApi) StartMonkeyApi(c *gin.Context) {
	var startMonkeyReq monkeyReq.StartMonkeyReq
	_ = c.ShouldBindJSON(&startMonkeyReq)
	if err := global.Validate.Struct(&startMonkeyReq); err != nil {
		global.GVA_LOG.Error("参数缺失", zap.Error(err))
		response.FailWithMessage("参数缺失", c)
		return
	}
	response.OkWithMessage("成功发起monkey测试，稍后生成测试报告", c)

}
