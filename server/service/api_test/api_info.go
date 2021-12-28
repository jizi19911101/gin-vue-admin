package api_test

//
//import (
//	"github.com/jizi19911101/gin-vue-admin/server/global"
//	"github.com/jizi19911101/gin-vue-admin/server/model/api_test"
//	autoCodeReq "github.com/jizi19911101/gin-vue-admin/server/model/autocode/request"
//	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
//)
//
//type ApiInfoService struct {
//}
//
//// CreateApiInfo 创建ApiInfo记录
//// Author [piexlmax](https://github.com/piexlmax)
//func (apiInfoService *ApiInfoService) CreateApiInfo(apiInfo api_test.ApiInfo) (err error) {
//	err = global.GVA_DB.Create(&apiInfo).Error
//	return err
//}
//
//// DeleteApiInfo 删除ApiInfo记录
//// Author [piexlmax](https://github.com/piexlmax)
//func (apiInfoService *ApiInfoService) DeleteApiInfo(apiInfo api_test.ApiInfo) (err error) {
//	err = global.GVA_DB.Delete(&apiInfo).Error
//	return err
//}
//
//// DeleteApiInfoByIds 批量删除ApiInfo记录
//// Author [piexlmax](https://github.com/piexlmax)
//func (apiInfoService *ApiInfoService) DeleteApiInfoByIds(ids request.IdsReq) (err error) {
//	err = global.GVA_DB.Delete(&[]api_test.ApiInfo{}, "id in ?", ids.Ids).Error
//	return err
//}
//
//// UpdateApiInfo 更新ApiInfo记录
//// Author [piexlmax](https://github.com/piexlmax)
//func (apiInfoService *ApiInfoService) UpdateApiInfo(apiInfo api_test.ApiInfo) (err error) {
//	//err = global.GVA_DB.Updates(&apiInfo).Error
//	err = global.GVA_DB.Model(apiInfo).Select("*").Omit("CreatedAt").Updates(&apiInfo).Error
//	return err
//}
//
//// GetApiInfo 根据id获取ApiInfo记录
//// Author [piexlmax](https://github.com/piexlmax)
//func (apiInfoService *ApiInfoService) GetApiInfo(id uint) (err error, apiInfo api_test.ApiInfo) {
//	err = global.GVA_DB.Where("id = ?", id).First(&apiInfo).Error
//	return
//}
//
//// GetApiInfoInfoList 分页获取ApiInfo记录
//// Author [piexlmax](https://github.com/piexlmax)
//func (apiInfoService *ApiInfoService) GetApiInfoInfoList(info autoCodeReq.ApiInfoSearch) (err error, list interface{}, total int64) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	// 创建db
//	db := global.GVA_DB.Model(&api_test.ApiInfo{})
//	var apiInfos []api_test.ApiInfo
//	// 如果有条件搜索 下方会自动创建搜索语句
//	if info.Name != "" {
//		db = db.Where("name LIKE ?", "%"+info.Name+"%")
//	}
//	if info.Method != "" {
//		db = db.Where("method = ?", info.Method)
//	}
//	if info.Url != "" {
//		db = db.Where("url LIKE ?", "%"+info.Url+"%")
//	}
//	if info.Params != "" {
//		db = db.Where("params = ?", info.Params)
//	}
//	if info.Project != "" {
//		db = db.Where("project LIKE ?", "%"+info.Project+"%")
//	}
//	if info.Module != "" {
//		db = db.Where("module LIKE ?", "%"+info.Module+"%")
//	}
//	err = db.Count(&total).Error
//	if err != nil {
//		return
//	}
//	err = db.Limit(limit).Offset(offset).Find(&apiInfos).Error
//	return err, apiInfos, total
//}
