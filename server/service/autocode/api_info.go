package autocode

import (
	"github/jizi19911101/gin-vue-admin/server/global"
	"github/jizi19911101/gin-vue-admin/server/model/autocode"
	"github/jizi19911101/gin-vue-admin/server/model/common/request"
    autoCodeReq "github/jizi19911101/gin-vue-admin/server/model/autocode/request"
)

type ApiInfoService struct {
}

// CreateApiInfo 创建ApiInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiInfoService *ApiInfoService) CreateApiInfo(apiInfo autocode.ApiInfo) (err error) {
	err = global.GVA_DB.Create(&apiInfo).Error
	return err
}

// DeleteApiInfo 删除ApiInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiInfoService *ApiInfoService)DeleteApiInfo(apiInfo autocode.ApiInfo) (err error) {
	err = global.GVA_DB.Delete(&apiInfo).Error
	return err
}

// DeleteApiInfoByIds 批量删除ApiInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiInfoService *ApiInfoService)DeleteApiInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ApiInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateApiInfo 更新ApiInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiInfoService *ApiInfoService)UpdateApiInfo(apiInfo autocode.ApiInfo) (err error) {
	err = global.GVA_DB.Save(&apiInfo).Error
	return err
}

// GetApiInfo 根据id获取ApiInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiInfoService *ApiInfoService)GetApiInfo(id uint) (err error, apiInfo autocode.ApiInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&apiInfo).Error
	return
}

// GetApiInfoInfoList 分页获取ApiInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiInfoService *ApiInfoService)GetApiInfoInfoList(info autoCodeReq.ApiInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.ApiInfo{})
    var apiInfos []autocode.ApiInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Method != "" {
        db = db.Where("method = ?",info.Method)
    }
    if info.Url != "" {
        db = db.Where("url LIKE ?","%"+ info.Url+"%")
    }
    if info.Params != "" {
        db = db.Where("params = ?",info.Params)
    }
    if info.Project != "" {
        db = db.Where("project LIKE ?","%"+ info.Project+"%")
    }
    if info.Module != "" {
        db = db.Where("module LIKE ?","%"+ info.Module+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&apiInfos).Error
	return err, apiInfos, total
}
