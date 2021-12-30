package api_test

import (
	"io/ioutil"
	"os"
	"regexp"

	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/api_test"

	"github.com/jizi19911101/gin-vue-admin/server/utils"
)

type ApiTestcaseService struct {
}

// ApiTestcaseCode 拉取接口自动化代码
func (apiTestcaseService *ApiTestcaseService) ApiTestcaseCode() (err error) {
	folder, _ := os.Getwd()
	targetFolder := folder + "/apiTestcaseCode"
	if err, _ := os.Stat(targetFolder); err != nil {
		os.Mkdir(targetFolder, 755)
	}

	err = utils.OsExecClone(targetFolder, "https://git-ext.chumanapp.com/chuman-test/chuman-api-test-new")
	return
}

// ApiTestcaseCode 解析接口自动化代码模块
func (apiTestcaseService *ApiTestcaseService) ParseApiTestcaseModule() (err error) {
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
		if fileName != "__init__.py" {
			moduleInfoListMap[fileName] = api_test.ModuleInfo{
				Name: fileName,
			}
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

// ApiTestcaseCode 解析接口自动化代码接口
func (apiTestcaseService *ApiTestcaseService) ParseApiTestcaseApi() (err error) {
	//取出模块
	moduleList := make([]api_test.ModuleInfo, 0)
	db := global.GVA_DB.Model(&api_test.ModuleInfo{})
	db.Find(&moduleList)

	//模块为 0，结束
	if len(moduleList) == 0 {
		return
	}

	//模块不为0，解析文件
	for _, v := range moduleList {
		folder, _ := os.Getwd()
		targetFolder := folder + "/apiTestcaseCode/testcases/" + v.Name
		targetFileList := make([]string, 0)
		if err, _ := os.Stat(targetFolder); err != nil {
			fileInfoList, _ := ioutil.ReadDir(targetFolder)
			for i := range fileInfoList {
				reg := regexp.MustCompile(`test_(.*?)\.py`)
				targetFile := reg.FindStringSubmatch(fileInfoList[i].Name())
				if len(targetFile) != 0 {
					targetFileList = append(targetFileList, targetFile[1])
				}
			}
		} else {
			global.GVA_LOG.Error("解析接口自动化代码接口出错")
		}
		if len(targetFileList) != 0 {
			db := global.GVA_DB.Model(&api_test.ApiInfo{})
			apiList := make([]api_test.ApiInfo, 0)
			for _, value := range targetFileList {
				apiList = append(apiList, api_test.ApiInfo{
					Name:         value,
					Module:       v.Name,
					Organization: "触漫",
				})
			}
			db.Create(&apiList)
		}

	}
	return
}
