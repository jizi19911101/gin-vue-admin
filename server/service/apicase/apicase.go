package apicase

import (
	"os/exec"

	"github.com/jizi19911101/gin-vue-admin/server/global"

	"github.com/jizi19911101/gin-vue-admin/server/model/apicase"
	apicaseReq "github.com/jizi19911101/gin-vue-admin/server/model/apicase/request"
)

type ApiCaseService struct {
}

func (apiCaseService *ApiCaseService) RunApiCase(runApiCaseReq apicaseReq.RunApiCaseReq) error {
	user := "huangweinan"
	userToken := "11b6fc0ee0647ba08c638bf188da4b4c2c"
	env := runApiCaseReq.Env
	module := runApiCaseReq.Module
	api := runApiCaseReq.Api
	caseName := runApiCaseReq.Case
	url := "http://jk-dev.chumanyun.com/job/qa-p0接口自动化测试/buildWithParameters"
	data := "envir=" + env
	if len(module) != 0 {
		data = data + "&module=" + module
	}
	if len(api) != 0 {
		data = data + "&api=" + "test_" + api + ".py"
	}
	if len(caseName) != 0 {
		var testcase = &apicase.ApiTestcase{}
		db := global.GVA_DB.Model(&apicase.ApiTestcase{})
		db.Select("class").Where("name = ?", caseName).Find(&testcase)
		if len(testcase.Class) != 0 {
			data = data + "::" + testcase.Class + "::" + caseName
		}

	}
	cmd := exec.Command("curl", url, "--user", user+":"+userToken, "--data", data)
	out, err := cmd.CombinedOutput()
	global.GVA_LOG.Debug(string(out) + "outoutoutout22")
	if err != nil {
		return err
	}
	return nil
}
