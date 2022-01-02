package organization

import (
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/organization"
	organizationReq "github.com/jizi19911101/gin-vue-admin/server/model/organization/request"
)

type EnvConfigService struct {
}

// CreateEnvConfig 创建EnvConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (envConfigService *EnvConfigService) CreateEnvConfig(envConfig organization.EnvConfig) (err error) {
	err = global.GVA_DB.Create(&envConfig).Error
	return err
}

// DeleteEnvConfig 删除EnvConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (envConfigService *EnvConfigService) DeleteEnvConfig(envConfig organization.EnvConfig) (err error) {
	err = global.GVA_DB.Delete(&envConfig).Error
	return err
}

// DeleteEnvConfigByIds 批量删除EnvConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (envConfigService *EnvConfigService) DeleteEnvConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]organization.EnvConfig{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEnvConfig 更新EnvConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (envConfigService *EnvConfigService) UpdateEnvConfig(envConfig organization.EnvConfig) (err error) {
	//fmt.Println(envConfig, "8989898989")
	//err = global.GVA_DB.Save(&envConfig).Error
	err = global.GVA_DB.Select("*").Omit("created_at").Updates(&envConfig).Error
	return err
}

// GetEnvConfig 根据id获取EnvConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (envConfigService *EnvConfigService) GetEnvConfig(id uint) (err error, envConfig organization.EnvConfig) {
	err = global.GVA_DB.Where("id = ?", id).First(&envConfig).Error
	return
}

// GetEnvConfigInfoList 分页获取EnvConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (envConfigService *EnvConfigService) GetEnvConfigInfoList(info organizationReq.EnvConfigSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&organization.EnvConfig{})
	var envConfigs []organization.EnvConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Base_url != "" {
		db = db.Where("base_url LIKE ?", "%"+info.Base_url+"%")
	}
	//if info.Project != "" {
	//	db = db.Where("organization LIKE ?", "%"+info.Project+"%")
	//}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&envConfigs).Error
	return err, envConfigs, total
}
