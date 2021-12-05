package api_test

import (
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/api_test"
	autoCodeReq "github.com/jizi19911101/gin-vue-admin/server/model/autocode/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
)

type ModuleService struct {
}

// CreateModule 创建Module记录
// Author [piexlmax](https://github.com/piexlmax)
func (moduleService *ModuleService) CreateModule(module api_test.Module) (err error) {
	err = global.GVA_DB.Create(&module).Error
	return err
}

// DeleteModule 删除Module记录
// Author [piexlmax](https://github.com/piexlmax)
func (moduleService *ModuleService) DeleteModule(module api_test.Module) (err error) {
	err = global.GVA_DB.Delete(&module).Error
	return err
}

// DeleteModuleByIds 批量删除Module记录
// Author [piexlmax](https://github.com/piexlmax)
func (moduleService *ModuleService) DeleteModuleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]api_test.Module{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateModule 更新Module记录
// Author [piexlmax](https://github.com/piexlmax)
func (moduleService *ModuleService) UpdateModule(module api_test.Module) (err error) {
	err = global.GVA_DB.Save(&module).Error
	return err
}

// GetModule 根据id获取Module记录
// Author [piexlmax](https://github.com/piexlmax)
func (moduleService *ModuleService) GetModule(id uint) (err error, module api_test.Module) {
	err = global.GVA_DB.Where("id = ?", id).First(&module).Error
	return
}

// GetModuleInfoList 分页获取Module记录
// Author [piexlmax](https://github.com/piexlmax)
func (moduleService *ModuleService) GetModuleInfoList(info autoCodeReq.ModuleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&api_test.Module{})
	var modules []api_test.Module
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Project != "" {
		db = db.Where("project LIKE ?", "%"+info.Project+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&modules).Error
	return err, modules, total
}
