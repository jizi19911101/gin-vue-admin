package apicase

import (
	"os/exec"

	"github.com/jizi19911101/gin-vue-admin/server/global"

	apicaseReq "github.com/jizi19911101/gin-vue-admin/server/model/apicase/request"
)

type ApiCaseService struct {
}

func (apiCaseService *ApiCaseService) RunApiCase(runApiCaseReq apicaseReq.RunApiCaseReq) error {
	user := "huangweinan"
	userToken := "11b6fc0ee0647ba08c638bf188da4b4c2c"
	env := runApiCaseReq.Env
	url := "http://jk-dev.chumanyun.com/job/qa-p0接口自动化测试/buildWithParameters"
	cmd := exec.Command("curl", url, "--user", user+":"+userToken, "--data", "envir="+env)
	out, err := cmd.CombinedOutput()
	global.GVA_LOG.Debug(string(out) + "outoutoutout22")
	if err != nil {
		return err
	}
	return nil
}
