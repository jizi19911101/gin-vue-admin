package api_test

import (
	"io/ioutil"
	"os"

	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/api_test"

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
	var resModuleInfoList []api_test.ModuleInfo
	delModuleInfoList := make([]uint, 0)
	var count int64

	// 解析出模块
	pwd, _ := os.Getwd()
	fileInfoList, err := ioutil.ReadDir(pwd + "/apiTestcaseCode/testcases")
	if len(fileInfoList) == 0 {
		return
	}
	moduleInfoList := make([]api_test.ModuleInfo, 0)
	moduleInfoListMap := make(map[string]api_test.ModuleInfo)
	for i := range fileInfoList {
		fileName := fileInfoList[i].Name()
		moduleInfoListMap[fileName] = api_test.ModuleInfo{
			Name: fileName,
		}
	}

	// 查库查出模块
	db := global.GVA_DB.Model(&api_test.ModuleInfo{})
	db.Find(&resModuleInfoList).Count(&count)

	// 把增量模块插入库
	if count == 0 {
		for _, v := range moduleInfoListMap {
			moduleInfoList = append(moduleInfoList, v)
		}
		db.Create(&moduleInfoList)
		return
	} else {
		for _, v := range resModuleInfoList {
			_, ok := moduleInfoListMap[v.Name]
			if ok {
				delete(moduleInfoListMap, v.Name)
			} else {
				delModuleInfoList = append(delModuleInfoList, v.ID)
			}
		}
	}

	if len(moduleInfoListMap) != 0 {
		for _, v := range moduleInfoListMap {
			moduleInfoList = append(moduleInfoList, v)
		}
		db.Create(&moduleInfoList)
	}

	if len(delModuleInfoList) != 0 {
		db.Delete(&api_test.ModuleInfo{}, delModuleInfoList)
	}

	return
}
