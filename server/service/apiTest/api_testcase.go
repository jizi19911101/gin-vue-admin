package apiTest

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"go.uber.org/zap"

	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/apiTest"
	"github.com/jizi19911101/gin-vue-admin/server/utils"
)

type ApiTestcaseService struct {
}

// ApiTestcaseCode 拉取接口自动化代码
func (apiTestcaseService *ApiTestcaseService) SyncApiTestCase() (err error) {
	tmpDir, err := ioutil.TempDir("./", "temp_*")
	defer os.RemoveAll(tmpDir)
	if err != nil {
		return err
	}
	err = utils.OsExecClone(tmpDir, "https://git-ext.chumanapp.com/chuman-test/chuman-api-test-new")
	if err != nil {
		return err
	}
	err = apiTestcaseService.ParseApiTestcaseModule(tmpDir)
	if err != nil {
		return err
	}
	err = apiTestcaseService.ParseApiTestcaseApi(tmpDir)
	if err != nil {
		return err
	}

	err = apiTestcaseService.ParseApiTestcase(tmpDir)
	if err != nil {
		return err
	}
	return
}

// ApiTestcaseCode 解析接口自动化代码模块
func (apiTestcaseService *ApiTestcaseService) ParseApiTestcaseModule(tmpDir string) (err error) {
	// 解析出模块
	fileInfoList, err := ioutil.ReadDir(tmpDir + "/testcases")
	if len(fileInfoList) == 0 {
		return
	}

	parseModuleMap := make(map[string]apiTest.Module)
	for i := range fileInfoList {
		fileName := fileInfoList[i].Name()
		if fileName != "__init__.py" {
			parseModuleMap[fileName] = apiTest.Module{
				Name: fileName,
			}
		}
	}

	// 查库查出模块
	db := global.GVA_DB.Model(&apiTest.Module{})

	var moduleList []apiTest.Module
	var count int64
	db.Find(&moduleList).Count(&count)

	// 把增量模块插入库
	if count == 0 {
		list := make([]apiTest.Module, 0)
		for _, module := range parseModuleMap {
			list = append(list, module)
		}
		db.Create(&list)
		return
	}

	moduleMap := make(map[string]apiTest.Module, 0)
	for _, module := range moduleList {
		moduleMap[module.Name] = module
	}

	addModuleInfoList := make([]apiTest.Module, 0)
	delModuleInfoList := make([]uint, 0)

	for _, module := range moduleList {
		if _, ok := parseModuleMap[module.Name]; !ok {
			delModuleInfoList = append(delModuleInfoList, module.ID)
		}
	}

	for key, module := range parseModuleMap {
		if _, ok := moduleMap[key]; !ok {
			addModuleInfoList = append(addModuleInfoList, module)
		}
	}

	if len(addModuleInfoList) != 0 {
		db.Create(&addModuleInfoList)
	}

	if len(delModuleInfoList) != 0 {
		db.Delete(&apiTest.Module{}, delModuleInfoList)
	}

	return
}

// ApiTestcaseCode 解析接口自动化代码接口
func (apiTestcaseService *ApiTestcaseService) ParseApiTestcaseApi(tmpDir string) error {
	//取出模块
	moduleList := make([]apiTest.Module, 0)
	db := global.GVA_DB.Model(&apiTest.Module{})
	db.Find(&moduleList)

	//模块为 0，结束
	if len(moduleList) == 0 {
		return nil
	}

	//模块不为0，解析每个模块下的接口文件
	for _, module := range moduleList {
		moduleFolder := tmpDir + "/testcases/" + module.Name
		if _, err := os.Stat(moduleFolder); err != nil {
			global.GVA_LOG.Error("解析模块目录出错", zap.Error(err))
			return err
		}

		parseApiList := make([]string, 0)
		list, _ := ioutil.ReadDir(moduleFolder)
		for i := range list {
			reg := regexp.MustCompile(`test_(.*?)\.py`)
			result := reg.FindStringSubmatch(list[i].Name())
			if len(result) != 0 {
				parseApiList = append(parseApiList, result[1])
			}
		}

		// 接口文件不为 0，就进行处理并存入数据库
		if len(parseApiList) != 0 {
			parseApiMap := make(map[string]apiTest.Api, 0)
			for _, a := range parseApiList {
				parseApiMap[a] = apiTest.Api{
					Name:   a,
					Module: module.Name,
				}
			}

			//查出该模块下的接口数据
			apiList := make([]apiTest.Api, 0)
			var count int64
			db := global.GVA_DB.Model(&apiTest.Api{})
			db.Where("module = ?", module.Name).Find(&apiList).Count(&count)

			// 该模块下的接口数据为0，直接插入
			if count == 0 {
				apiList := make([]apiTest.Api, 0)
				for _, api := range parseApiMap {
					apiList = append(apiList, api)
				}
				db.Create(&apiList)
			}

			apiMap := make(map[string]apiTest.Api, 0)
			for _, a := range apiList {
				apiMap[a.Name] = a
			}

			// 该模块下的接口数据不为0，增量插入
			delApiList := make([]uint, 0)
			addApiList := make([]apiTest.Api, 0)

			for key, value := range apiMap {
				if _, ok := parseApiMap[key]; !ok {
					delApiList = append(delApiList, value.ID)
				}
			}

			for key, value := range parseApiMap {
				if _, ok := apiMap[key]; !ok {
					addApiList = append(addApiList, value)
				}
			}

			if len(addApiList) != 0 {
				db.Create(&addApiList)
			}

			if len(delApiList) != 0 {
				db.Delete(&apiTest.Api{}, delApiList)
			}

		}

	}
	return nil
}

func (apiTestcaseService *ApiTestcaseService) ParseApiTestcase(tmpDir string) (err error) {
	// 取出所有接口
	apiList := make([]apiTest.Api, 0)
	db := global.GVA_DB.Model(&apiTest.Api{})
	db.Find(&apiList)

	// 接口数量为0结束
	if len(apiList) == 0 {
		return
	}

	// 读接口文件
	for _, api := range apiList {
		targetFile := tmpDir + "/testcases/" + api.Module + "/test_" + api.Name + ".py"
		var className string
		caseList := make([]string, 0)
		if _, err := os.Stat(targetFile); err == nil {

			func() {
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
			}()

		} else {
			global.GVA_LOG.Error("解析接口自动化用例出错")
		}

		// 用例数为0，结束
		if len(caseList) == 0 {
			return
		}

		//用例数不为 0，读出数据库的用例
		stockTestcase := make([]apiTest.ApiTestcase, 0)
		if className != "" {
			db := global.GVA_DB.Model(&apiTest.ApiTestcase{})
			db.Where("module = ? AND api = ?", api.Module, api.Name).Find(&stockTestcase)
		} else {
			global.GVA_LOG.Error("接口自动化文件用例类名解析出错")
		}

		caseListMap := make(map[string]apiTest.ApiTestcase, 0)
		for _, v := range caseList {
			caseListMap[v] = apiTest.ApiTestcase{
				Name:   v,
				Module: api.Module,
				Api:    api.Name,
				Class:  className,
			}

		}

		//数据库用例数为0，直接加入
		db := global.GVA_DB.Model(&apiTest.ApiTestcase{})
		cases := make([]apiTest.ApiTestcase, 0)
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
			db.Delete(&apiTest.ApiTestcase{}, delCaseList)
		}

	}

	return
}
