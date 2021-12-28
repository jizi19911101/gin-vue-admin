package api_test

import "github.com/jizi19911101/gin-vue-admin/server/utils"

type ApiTestcaseService struct {
}

func (apiTestcaseService *ApiTestcaseService) ApiTestcaseCode() (err error) {
	err = utils.OsExecClone("  ", "https://git-ext.chumanapp.com/chuman-test/chuman-api-test-new")
	return
}
