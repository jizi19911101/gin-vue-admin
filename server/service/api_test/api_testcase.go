package api_test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

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
	if _, err := os.Stat(targetFolder); err != nil {
		os.Mkdir(targetFolder, 755)
	} else {
		os.Remove(targetFolder)
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
		for _, module := range moduleInfoListMap {
			moduleInfoList = append(moduleInfoList, module)
		}
		db.Create(&moduleInfoList)
		return
	} else {
		for _, module := range resModuleInfoList {
			_, ok := moduleInfoListMap[module.Name]
			if ok {
				delete(moduleInfoListMap, module.Name)
			} else {
				delModuleInfoList = append(delModuleInfoList, module.ID)
			}
		}
	}

	if len(moduleInfoListMap) != 0 {
		for _, module := range moduleInfoListMap {
			moduleInfoList = append(moduleInfoList, module)
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

	//模块不为0，解析每个模块下的接口文件
	for _, module := range moduleList {
		folder, _ := os.Getwd()
		targetFolder := folder + "/apiTestcaseCode/testcases/" + module.Name
		targetFileList := make([]string, 0)
		if _, err := os.Stat(targetFolder); err == nil {
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
		// 接口文件不为 0，就存入数据库
		if len(targetFileList) != 0 {
			db := global.GVA_DB.Model(&api_test.ApiInfo{})
			apiList := make([]api_test.ApiInfo, 0)
			apiListMap := make(map[string]api_test.ApiInfo, 0)
			for _, apiName := range targetFileList {
				apiListMap[apiName] = api_test.ApiInfo{
					Name:         apiName,
					Module:       module.Name,
					Organization: "触漫",
				}
			}

			//查出该模块下的接口数据
			resApiList := make([]api_test.ApiInfo, 0)
			var count int64
			db.Where("module = ?", module.Name).Find(&resApiList).Count(&count)

			// 该模块下的接口数据为0，直接插入
			if count == 0 {
				for _, api := range apiListMap {
					apiList = append(apiList, api)
				}
				db.Create(&apiList)
			}

			// 该模块下的接口数据不为0，增量插入
			delApiList := make([]uint, 0)
			for _, api := range resApiList {
				_, ok := apiListMap[api.Name]
				if ok {
					delete(apiListMap, api.Name)
				} else {
					delApiList = append(delApiList, api.ID)
				}
			}

			if len(apiListMap) != 0 {
				for _, api := range apiListMap {
					apiList = append(apiList, api)
				}
				db.Create(&apiList)
			}
			if len(delApiList) != 0 {
				fmt.Println(delApiList, "delApiList")
				db.Delete(&api_test.ApiInfo{}, delApiList)
			}

		}

	}
	return
}

func (apiTestcaseService *ApiTestcaseService) ParseApiTestcase() (err error) {
	// 取出所有接口
	apiList := make([]api_test.ApiInfo, 0)
	db := global.GVA_DB.Model(&api_test.ApiInfo{})
	db.Find(&apiList)

	// 接口数量为0结束
	if len(apiList) == 0 {
		return
	}

	// 读接口文件
	for _, api := range apiList {
		folder, _ := os.Getwd()
		targetFile := folder + "/apiTestcaseCode/testcases/" + api.Module + "/test_" + api.Name + ".py"
		var className string
		caseList := make([]string, 0)
		if _, err := os.Stat(targetFile); err == nil {
			//文件存在，解析出用例
			file, err := os.Open(targetFile)
			defer file.Close()
			if err != nil {
				panic(err)
			} else {
				scanner := bufio.NewScanner(file)

				for scanner.Scan() {
					regClass := regexp.MustCompile("class(.*?):")
					caseClass := regClass.FindStringSubmatch(scanner.Text())

					regCase := regexp.MustCompile(`^def(.*?)\(`)
					caseName := regCase.FindStringSubmatch(strings.TrimSpace(scanner.Text()))

					if len(caseClass) != 0 {
						className = caseClass[1]
					}

					if len(caseName) != 0 {
						caseList = append(caseList, caseName[1])
					}

				}
			}
		} else {
			global.GVA_LOG.Error("解析接口自动化用例出错")
		}

		// 用例数为0，结束
		if len(caseList) == 0 {
			return
		}

		//用例数不为 0，读出数据库的用例
		stockTestcase := make([]api_test.ApiTestcase, 0)
		if className != "" {
			db := global.GVA_DB.Model(&api_test.ApiTestcase{})
			db.Where("module = ? AND api = ?", api.Module, api.Name).Find(&stockTestcase)
		} else {
			global.GVA_LOG.Error("接口自动化文件用例类名解析出错")
		}

		caseListMap := make(map[string]api_test.ApiTestcase, 0)
		for _, v := range caseList {
			caseListMap[v] = api_test.ApiTestcase{
				Name:   v,
				Module: api.Module,
				Api:    api.Name,
				Class:  className,
			}

		}

		//数据库用例数为0，直接加入
		db := global.GVA_DB.Model(&api_test.ApiTestcase{})
		cases := make([]api_test.ApiTestcase, 0)
		if len(stockTestcase) == 0 {
			for _, v := range caseListMap {
				cases = append(cases, v)
			}
			db.Create(&cases)
		}

		//数据库用例数不为0，进行筛选再加到数据库
		delCaseList := make([]uint, 0)
		for _, c := range stockTestcase {
			_, ok := caseListMap[c.Name]
			if ok {
				delete(caseListMap, c.Name)
			} else {
				delCaseList = append(delCaseList, c.ID)
			}
		}

		if len(caseListMap) != 0 {
			for _, v := range caseListMap {
				cases = append(cases, v)
			}
			db.Create(&cases)
		}

		if len(delCaseList) != 0 {
			fmt.Println(delCaseList, "delCaseListdelCaseList")
			db.Delete(&api_test.ApiTestcase{}, delCaseList)
		}

	}

	return
}
