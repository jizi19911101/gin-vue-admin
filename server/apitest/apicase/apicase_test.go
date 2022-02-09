package apicase

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/jizi19911101/gin-vue-admin/server/initialize"
)

func TestApiCase(t *testing.T) {
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

	// 查询模块
	obj := e.GET("/apiCase/moduleList").
		WithHeader("x-token", Token).
		WithQuery("page", 1).
		WithQuery("pageSize", 10).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("获取模块列表成功")

	// 查询接口
	obj = e.GET("/apiCase/apiList").
		WithHeader("x-token", Token).
		WithQuery("page", 1).
		WithQuery("pageSize", 10).
		WithQuery("module", "adventuregame").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("获取接口列表成功")

	// 查询用例
	obj = e.GET("/apiCase/caseList").
		WithHeader("x-token", Token).
		WithQuery("page", 1).
		WithQuery("pageSize", 10).
		WithQuery("module", "adventuregame").
		WithQuery("api", "get_game_style_info").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("获取用例列表成功！")

	// 查询报告
	obj = e.GET("/apiCase/reportList").
		WithHeader("x-token", Token).
		WithQuery("page", 1).
		WithQuery("pageSize", 10).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("获取报告列表成功!")

}
