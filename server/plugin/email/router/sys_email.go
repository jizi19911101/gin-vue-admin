package router

import (
	"github.com/jizi19911101/gin-vue-admin/server/middleware"
	"github.com/jizi19911101/gin-vue-admin/server/plugin/email/api"
	"github.com/gin-gonic/gin"
)

type EmailRouter struct {
}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord())
	var EmailApi = api.ApiGroupApp.EmailApi.EmailTest
	var SendEmail = api.ApiGroupApp.EmailApi.SendEmail
	{
		emailRouter.POST("emailTest", EmailApi)  // 发送测试邮件
		emailRouter.POST("sendEmail", SendEmail) // 发送邮件
	}
}
