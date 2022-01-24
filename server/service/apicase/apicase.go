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
	user := global.GVA_CONFIG.Jenkins.User
	token := global.GVA_CONFIG.Jenkins.Token
	global.GVA_LOG.Debug("jenkins用户：" + user + token)
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
		var testcase = &apicase.ApiCase{}
		db := global.GVA_DB.Model(&apicase.ApiCase{})
		db.Select("class").Where("name = ? AND api = ? ", caseName, api).Find(&testcase)
		if len(testcase.Class) != 0 {
			data = data + "&class=" + testcase.Class + "&case=" + caseName
		}

	}
	global.GVA_LOG.Debug("调接口自动化job的data参数：" + data)

	cmd := exec.Command("curl", url, "--user", user+":"+token, "--data", data)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

func (apiCaseService *ApiCaseService) ModuleList(info apicaseReq.ModuleSearch) (error, interface{}, int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&apicase.Module{})
	var moduleList []apicase.Module
	var total int64

	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if err := db.Count(&total).Error; err != nil {
		return err, nil, 0
	}

	err := db.Limit(limit).Offset(offset).Find(&moduleList).Error
	return err, moduleList, total

}

func (apiCaseService *ApiCaseService) ApiList(info apicaseReq.ApiSearch) (error, interface{}, int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&apicase.Api{})
	var apiList []apicase.Api
	var total int64

	if info.Name != "" {
		db.Where("name = ?", info.Name)
	}

	if info.Module != "" {
		db.Where("module = ?", info.Module)
	}

	if err := db.Count(&total).Error; err != nil {
		return err, nil, 0
	}
	err := db.Limit(limit).Offset(offset).Find(&apiList).Error
	return err, apiList, total
}

func (apiCaseService *ApiCaseService) ApiCaseList(info apicaseReq.ApiCaseSearch) (error, interface{}, int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&apicase.ApiCase{})
	var total int64
	var apiCaseList []apicase.ApiCase

	if info.Module != "" {
		db.Where("module = ?", info.Module)
	}
	if info.Api != "" {
		db.Where("api = ? ", info.Api)
	}
	if info.Name != "" {
		db.Where("name = ? ", info.Name)

	}

	if err := db.Count(&total).Error; err != nil {
		return err, nil, 0
	}
	err := db.Limit(limit).Offset(offset).Find(&apiCaseList).Error

	return err, apiCaseList, total
}

func (apiCaseService *ApiCaseService) ReportList(info apicaseReq.ReportSearch) (error, interface{}, int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&apicase.Report{})
	var total int64
	reportList := make([]apicase.Report, 0)

	if info.Name != "" {
		db.Where("name = ?", info.Name)
	}

	if err := db.Count(&total).Error; err != nil {
		return err, nil, 0
	}
	err := db.Limit(limit).Offset(offset).Find(&reportList).Error
	return err, reportList, total

}
