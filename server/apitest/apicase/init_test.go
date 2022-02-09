package apicase

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/jizi19911101/gin-vue-admin/server/initialize"
)

var Token = ""

func init() {
	t := &testing.T{}
	handler := initialize.Routers()

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	userInfo := map[string]interface{}{
		"username":  "admin",
		"password":  "123456",
		"captcha":   "123123",
		"captchaId": "HZyBBHFXCq8TUGzXD1Fa",
	}
	obj := e.POST("/base/login").
		WithJSON(userInfo).
		Expect().
		Status(http.StatusOK).JSON().Object()

	data := obj.Raw()["data"]
	Token = data.(map[string]interface{})["token"].(string)

}
