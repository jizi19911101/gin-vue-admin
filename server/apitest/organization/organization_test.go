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
	// 创建3个
	organization := map[string]string{
		"name": "单元测试",
	}
	for i := 0; i < 3; i++ {
		obj := e.POST("/organization/createOrganization").
			WithHeader("x-token", Token).
			WithJSON(organization).
			Expect().
			Status(http.StatusOK).JSON().Object()

		obj.Value("msg").Equal("创建成功")
	}

	// 查询
	obj := e.GET("/organization/getOrganizationList").
		WithHeader("x-token", Token).
		WithQuery("name", "单元测试").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.Value("msg").Equal("获取成功")

	data := obj.Raw()["data"].(map[string]interface{})
	list := data["list"].([]interface{})
	idList := make([]float64, 0, 10)
	for _, i := range list {
		id := i.(map[string]interface{})["ID"].(float64)
		idList = append(idList, id)

	}

	//多个删除
	delIds := map[string][]float64{
		"ids": idList,
	}
	obj = e.DELETE("/organization/deleteOrganizationByIds").
		WithHeader("x-token", Token).
		WithJSON(delIds).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("批量删除成功")
}

func TestUpdateOrganization(t *testing.T) {
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
	// 创建
	organization := map[string]string{
		"name": "单元测试（修改）",
	}
	obj := e.POST("/organization/createOrganization").
		WithHeader("x-token", Token).
		WithJSON(organization).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("创建成功")

	// 查询id
	obj = e.GET("/organization/getOrganizationList").
		WithHeader("x-token", Token).
		WithQuery("name", "单元测试（修改）").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("获取成功")

	data := obj.Raw()["data"].(map[string]interface{})
	list := data["list"].([]interface{})
	id := list[0].(map[string]interface{})["ID"].(float64)

	// 修改
	updateConent := map[string]interface{}{
		"ID":   id,
		"name": "单元测试（修改2）",
	}
	obj = e.PUT("/organization/updateOrganization").
		WithHeader("x-token", Token).
		WithJSON(updateConent).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("更新成功")

	// 根据id查询
	obj = e.GET("/organization/findOrganization").
		WithHeader("x-token", Token).
		WithQuery("ID", id).
		Expect().
		Status(http.StatusOK).JSON().Object()
	//obj.Value("data").Object().Value("organization").Object().ValueEqual("name", "单元测试（修改2）")
	obj.Path("$.data.organization.name").String().Equal("单元测试（修改2）")

	// 删除单个
	delId := map[string]float64{
		"ID": id,
	}
	obj = e.DELETE("/organization/deleteOrganization").
		WithHeader("x-token", Token).
		WithJSON(delId).
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("删除成功")
}
