package sync

import (
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"

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
	//engine := gin.New()
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

	e.GET("/sync/syncApiTestcase").
		Expect().
		Status(http.StatusOK)

}
