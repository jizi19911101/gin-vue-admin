package autocode

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	apiInfoReq "github.com/jizi19911101/gin-vue-admin/server/model/apiInfo/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/jizi19911101/gin-vue-admin/server/model/autocode/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/request"
	"github.com/jizi19911101/gin-vue-admin/server/model/common/response"
	"github.com/jizi19911101/gin-vue-admin/server/service"
	"go.uber.org/zap"
	"strings"
)

type ApiInfoApi struct {
}

var apiInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.ApiInfoService


// CreateApiInfo 创建ApiInfo
// @Tags ApiInfo
// @Summary 创建ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ApiInfo true "创建ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiInfo/createApiInfo [post]
func (apiInfoApi *ApiInfoApi) CreateApiInfo(c *gin.Context) {
	var reqApiInfo apiInfoReq.ApiInfo
	_ = c.ShouldBindJSON(&reqApiInfo)
	apiInfo := autocode.ApiInfo{
		Name: reqApiInfo.Name,
		Url: reqApiInfo.Url,
		Method: reqApiInfo.Method,
		Project: reqApiInfo.Project,
		Module: reqApiInfo.Module,
		Params: strings.Join(reqApiInfo.Params,","),
	}
	//s2 := strings.Join(s1,",")
	if err := apiInfoService.CreateApiInfo(apiInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApiInfo 删除ApiInfo
// @Tags ApiInfo
// @Summary 删除ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ApiInfo true "删除ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiInfo/deleteApiInfo [delete]
func (apiInfoApi *ApiInfoApi) DeleteApiInfo(c *gin.Context) {
	var apiInfo autocode.ApiInfo
	_ = c.ShouldBindJSON(&apiInfo)
	if err := apiInfoService.DeleteApiInfo(apiInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteApiInfoByIds 批量删除ApiInfo
// @Tags ApiInfo
// @Summary 批量删除ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /apiInfo/deleteApiInfoByIds [delete]
func (apiInfoApi *ApiInfoApi) DeleteApiInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := apiInfoService.DeleteApiInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateApiInfo 更新ApiInfo
// @Tags ApiInfo
// @Summary 更新ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ApiInfo true "更新ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apiInfo/updateApiInfo [put]
func (apiInfoApi *ApiInfoApi) UpdateApiInfo(c *gin.Context) {
	var apiInfo autocode.ApiInfo
	_ = c.ShouldBindJSON(&apiInfo)
	if err := apiInfoService.UpdateApiInfo(apiInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindApiInfo 用id查询ApiInfo
// @Tags ApiInfo
// @Summary 用id查询ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ApiInfo true "用id查询ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apiInfo/findApiInfo [get]
func (apiInfoApi *ApiInfoApi) FindApiInfo(c *gin.Context) {
	var apiInfo autocode.ApiInfo
	_ = c.ShouldBindQuery(&apiInfo)
	if err, reapiInfo := apiInfoService.GetApiInfo(apiInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapiInfo": reapiInfo}, c)
	}
}

// GetApiInfoList 分页获取ApiInfo列表
// @Tags ApiInfo
// @Summary 分页获取ApiInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ApiInfoSearch true "分页获取ApiInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiInfo/getApiInfoList [get]
func (apiInfoApi *ApiInfoApi) GetApiInfoList(c *gin.Context) {
	var pageInfo autocodeReq.ApiInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := apiInfoService.GetApiInfoInfoList(pageInfo); err != nil {
		fmt.Println(err,"4325435435643")
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {

    	apiList := list.([]autocode.ApiInfo)
		resApiInfoList := make([]resApiInfo, 0,len(apiList))

		for i,api := range apiList {
			resApiInfo := resApiInfo{
				Name:api.Name,
				Method:api.Method,
				Url:api.Url,
				Project:api.Project,
				Module:api.Module,
			}
			if err := json.Unmarshal([]byte(apiList[i].Params), &resApiInfo.Params);err!=nil{
				fmt.Println(err,"errr12121212121212")
				response.FailWithMessage("param 解析失败", c)
			}
			//if err := strings.Split(apiList[i].Params, ",");err != nil{
			//	response.FailWithMessage("param 解析失败", c)
			//}
			resApiInfoList = append(resApiInfoList, resApiInfo)
		}
        fmt.Println("8989898989898")
        response.OkWithDetailed(response.PageResult{
            List:     resApiInfoList,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

type resApiInfo struct {
	Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar;"`
	Method  string `json:"method" form:"method" gorm:"column:method;comment:;type:char;"`
	Url  string `json:"url" form:"url" gorm:"column:url;comment:;type:varchar;"`
	Params  []map[string]string `json:"params" form:"params" gorm:"column:params;comment:;type:varchar;"`
	Project  string `json:"project" form:"project" gorm:"column:project;comment:;type:char;"`
	Module  string `json:"module" form:"module" gorm:"column:module;comment:;type:varchar;"`
}


