package sync

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"go.uber.org/zap"

	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/apicase"
	"github.com/jizi19911101/gin-vue-admin/server/utils"
)

type SyncService struct {
}

// ApiTestcaseCode 拉取接口自动化代码
func (syncService *SyncService) SyncApiTestCase() (err error) {
	tmpDir, err := ioutil.TempDir("./", "temp_*")
	defer os.RemoveAll(tmpDir)
	if err != nil {
		return err
	}
	err = utils.OsExecClone(tmpDir, "git@git-ext.chumanapp.com:chuman-test/chuman-api-test-new.git")
	if err != nil {
		return err
	}
	err = syncService.ParseApiTestcaseModule(tmpDir)
	if err != nil {
		return err
	}
	err = syncService.ParseApiTestcaseApi(tmpDir)
	if err != nil {
		return err
	}

	err = syncService.ParseApiTestcase(tmpDir)
	if err != nil {
		return err
	}
	return
}

// ApiTestcaseCode 解析接口自动化代码模块
func (syncService *SyncService) ParseApiTestcaseModule(tmpDir string) (err error) {
	// 解析出模块
	fileInfoList, err := ioutil.ReadDir(tmpDir + "/testcases")
	if len(fileInfoList) == 0 {
		return
	}

	parseModuleMap := make(map[string]apicase.Module)
	for i := range fileInfoList {
		fileName := fileInfoList[i].Name()
		if fileName != "__init__.py" {
			parseModuleMap[fileName] = apicase.Module{
				Name: fileName,
			}
		}
	}

	// 查库查出模块
	db := global.GVA_DB.Model(&apicase.Module{})

	var moduleList []apicase.Module
	var count int64
	db.Find(&moduleList).Count(&count)

	// 把增量模块插入库
	if count == 0 {
		list := make([]apicase.Module, 0)
		for _, module := range parseModuleMap {
			list = append(list, module)
		}
		db.Create(&list)
		return
	}

	moduleMap := make(map[string]apicase.Module, 0)
	for _, module := range moduleList {
		moduleMap[module.Name] = module
	}

	addModuleInfoList := make([]apicase.Module, 0)
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
		db.Delete(&apicase.Module{}, delModuleInfoList)
	}

	return
}

// ApiTestcaseCode 解析接口自动化代码接口
func (syncService *SyncService) ParseApiTestcaseApi(tmpDir string) error {
	//取出模块
	moduleList := make([]apicase.Module, 0)
	db := global.GVA_DB.Model(&apicase.Module{})
	db.Find(&moduleList)

	//模块为 0，结束
	if len(moduleList) == 0 {
		return nil
	}

	//模块不为0，解析每个模块下的接口文件
	for _, module := range moduleList {
		moduleFolder := tmpDir + "/testcases/" + module.Name
		if _, err := os.Stat(moduleFolder); err != nil {
			global.GVA_LOG.Error("模块目录不存在", zap.Error(err))
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
			parseApiMap := make(map[string]apicase.Api, 0)
			for _, a := range parseApiList {
				parseApiMap[a] = apicase.Api{
					Name:   a,
					Module: module.Name,
				}
			}

			//查出该模块下的接口数据
			apiList := make([]apicase.Api, 0)
			var count int64
			db := global.GVA_DB.Model(&apicase.Api{})
			db.Where("module = ?", module.Name).Find(&apiList).Count(&count)

			// 该模块下的接口数据为0，直接插入
			if count == 0 {
				apiList := make([]apicase.Api, 0)
				for _, api := range parseApiMap {
					apiList = append(apiList, api)
				}
				db.Create(&apiList)
				continue
			}

			apiMap := make(map[string]apicase.Api, 0)
			for _, a := range apiList {
				apiMap[a.Name] = a
			}

			// 该模块下的接口数据不为0，增量插入
			delApiList := make([]uint, 0)
			addApiList := make([]apicase.Api, 0)

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
				db.Delete(&apicase.Api{}, delApiList)
			}

		}

	}
	return nil
}

func (syncService *SyncService) ParseApiTestcase(tmpDir string) error {
	// 取出所有接口
	apiList := make([]apicase.Api, 0)
	db := global.GVA_DB.Model(&apicase.Api{})
	db.Find(&apiList)

	// 接口数量为0结束
	if len(apiList) == 0 {
		return nil
	}

	// 读接口文件
	for _, api := range apiList {
		apiFile := tmpDir + "/testcases/" + api.Module + "/test_" + api.Name + ".py"
		if _, err := os.Stat(apiFile); err != nil {
			global.GVA_LOG.Error("接口文件不存在", zap.Error(err))
			return err
		}

		//文件存在，解析出用例
		className, parseCaseList, err := parseCase(apiFile)
		if err != nil {
			global.GVA_LOG.Error("解析接口用例出错", zap.Error(err))
			return err
		}

		// 用例数为0，结束
		if len(parseCaseList) == 0 {
			continue
		}

		parseCaseMap := make(map[string]apicase.ApiCase, 0)
		for _, v := range parseCaseList {
			parseCaseMap[v] = apicase.ApiCase{
				Name:   v,
				Module: api.Module,
				Api:    api.Name,
				Class:  className,
			}
		}

		//用例数不为 0，读出数据库的用例
		db := global.GVA_DB.Model(&apicase.ApiCase{})
		caseList := make([]apicase.ApiCase, 0)
		if className == "" {
			global.GVA_LOG.Error("接口用例类名解析出错")
			//抛出错误
			return errors.New("接口用例类名解析出错")
		}
		db.Where("module = ? AND api = ?", api.Module, api.Name).Find(&caseList)

		//数据库用例数为0，直接加入
		if len(caseList) == 0 {
			list := make([]apicase.ApiCase, 0)
			for _, v := range parseCaseMap {
				list = append(list, v)
			}
			db.Create(&list)
			continue
		}

		caseMap := make(map[string]apicase.ApiCase, 0)
		for _, c := range caseList {
			caseMap[c.Name] = c
		}

		//数据库用例数不为0，进行筛选再加到数据库
		delCaseList := make([]uint, 0)
		addCaseList := make([]apicase.ApiCase, 0)

		for _, t := range caseMap {
			if _, ok := parseCaseMap[t.Name]; !ok {
				delCaseList = append(delCaseList, t.ID)
			}
		}

		for _, t := range parseCaseMap {
			if _, ok := caseMap[t.Name]; !ok {
				addCaseList = append(addCaseList, t)
			}
		}

		if len(addCaseList) != 0 {
			db.Create(&addCaseList)
		}

		if len(delCaseList) != 0 {
			db.Delete(&apicase.ApiCase{}, delCaseList)
		}

	}

	return nil
}

func (syncService *SyncService) SyncApiTestReport(report apicase.Report) error {
	err := global.GVA_DB.Create(&report).Error
	return err
}

func parseCase(apiFile string) (string, []string, error) {
	var className string
	parseCaseList := make([]string, 0)

	file, err := os.Open(apiFile)
	if err != nil {
		global.GVA_LOG.Error("打开接口文件出错", zap.Error(err))
		return "", nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reg := regexp.MustCompile("class(.*?):")
		result := reg.FindStringSubmatch(scanner.Text())
		if len(result) != 0 {
			className = strings.Trim(result[1], " ")
		}

		reg = regexp.MustCompile(`^def(.*?)\(`)
		result = reg.FindStringSubmatch(strings.TrimSpace(scanner.Text()))
		if len(result) != 0 {
			parseCaseList = append(parseCaseList, strings.Trim(result[1], " "))
		}

	}
	return className, parseCaseList, nil
}
