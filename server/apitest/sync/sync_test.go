package sync

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"

	"github.com/jizi19911101/gin-vue-admin/server/initialize"
)

// assert example
// import "github.com/stretchr/testify/assert"
// assert := assert.New(t)
// assert.Equal(123, 123, "they should be equal")

// httpexpect
// import "github.com/gavv/httpexpect/v2"
//

//func TestParseCase(t *testing.T) {
//	dir, _ := ioutil.TempDir("./", "temp_*")
//	defer os.RemoveAll(dir)
//	syncService := SyncService{}
//	syncService.ParseApiTestcase(dir)
//
//}

func TestSyncApiTestCase(t *testing.T) {
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

	obj := e.GET("/sync/syncApiTestcase").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("同步并解析接口自动化代码成功！")

}

func TestSyncApiTestReport(t *testing.T) {
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

	obj := e.GET("/sync/syncApiTestReport").
		WithQuery("name", "百度").
		WithQuery("url", "https://www.baidu.com/").
		Expect().
		Status(http.StatusOK).JSON().Object()

	obj.Value("msg").Equal("同步接口测试报告成功！")

}
