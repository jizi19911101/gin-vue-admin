package api_test

import (
	"os"

	"github.com/jizi19911101/gin-vue-admin/server/utils"
)

type ApiTestcaseService struct {
}

func (apiTestcaseService *ApiTestcaseService) ApiTestcaseCode() (err error) {
	folder, _ := os.Getwd()
	targetFolder := folder + "/apiTestcaseCode"
	if err, _ := os.Stat(targetFolder); err != nil {
		os.Mkdir(targetFolder, 755)
	}

	err = utils.OsExecClone(targetFolder, "https://git-ext.chumanapp.com/chuman-test/chuman-api-test-new")
	return
}
