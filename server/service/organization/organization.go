package organization

import (
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/organization"
	organizationReq "github.com/jizi19911101/gin-vue-admin/server/model/organization/request"
)

type OrganizationService struct {
}

// CreateOrganization 创建Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) CreateOrganization(organization organization.Organization) (err error) {
	err = global.GVA_DB.Create(&organization).Error
	return err
}

// DeleteOrganization 删除Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) DeleteOrganization(organization organization.Organization) (err error) {
	err = global.GVA_DB.Delete(&organization).Error
	return err
}

// DeleteOrganizationByIds 批量删除Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) DeleteOrganizationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]organization.Organization{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOrganization 更新Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) UpdateOrganization(organization organization.Organization) (err error) {
	//err = global.GVA_DB.Save(&organization).Error
	err = global.GVA_DB.Select("*").Omit("created_at").Updates(&organization).Error
	return err
}

// GetOrganization 根据id获取Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) GetOrganization(id uint) (err error, organization organization.Organization) {
	err = global.GVA_DB.Where("id = ?", id).First(&organization).Error
	return
}

// GetOrganizationInfoList 分页获取Organization记录
// Author [piexlmax](https://github.com/piexlmax)
func (organizationService *OrganizationService) GetOrganizationInfoList(info organizationReq.OrganizationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&organization.Organization{})
	var organizations []organization.Organization
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&organizations).Error
	return err, organizations, total
}
