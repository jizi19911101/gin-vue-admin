package api_test

import (
	"fmt"
	"io/ioutil"
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

func (apiTestcaseService *ApiTestcaseService) ParseApiTestcaseApi() (err error) {
	pwd, _ := os.Getwd()
	fileInfoList, err := ioutil.ReadDir(pwd + "/apiTestcaseCode/testcases")
	for i := range fileInfoList {
		fmt.Println(fileInfoList[i].Name()) //打印当前文件或目录下的文件或目录名
	}
	return
}
