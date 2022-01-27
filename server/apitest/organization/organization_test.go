package organization

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"

	"github.com/jizi19911101/gin-vue-admin/server/initialize"
)

func TestCreateOrganization(t *testing.T) {
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

	organization := map[string]string{
		"name": "单元测试",
	}
	obj := e.POST("/organization/createOrganization").
		WithHeader("x-token", Token).
		WithJSON(organization).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("创建成功")

}
