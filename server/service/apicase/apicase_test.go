package apicase

import (
	"testing"

	apicaseReq "github.com/jizi19911101/gin-vue-admin/server/model/apicase/request"
)

func TestApiCase(t *testing.T) {
	runApiCaseReq := apicaseReq.RunApiCaseReq{}
	x := ApiCaseService{}
	x.RunApiCase(runApiCaseReq)
	//user := "huangweinan"
	//userToken := "11b6fc0ee0647ba08c638bf188da4b4c2c"
	//env := "demo"
	//url := "http://jk-dev.chumanyun.com/job/qa-p0接口自动化测试/buildWithParameters"
	////cmd := exec.Command("curl", url, "--user", user+":"+userToken, "--data", "envir="+env)
	//cmd := "curl " + url + "  --user  " + user + ":" + userToken + " --data " + "envir=" + env
	//fmt.Println(cmd)
}
