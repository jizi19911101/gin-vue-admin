import service from '@/utils/request'

// @Tags Module
// @Summary 分页获取Module列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Module列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiCase/moduleList [get]
export const getModuleList = (params) => {
    return service({
      url: '/apiCase/moduleList',
      method: 'get',
      params
    })
  }


// @Tags Report
// @Summary 分页获取Report列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Report列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiCase/reportList [get]
export const getReportList = (params) => {
    return service({
      url: '/apiCase/reportList',
      method: 'get',
      params
    })
  }


// @Tags ApiList
// @Summary 分页获取ApiList列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ApiList列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiCase/apiList [get]
export const getApiList = (params) => {
    return service({
      url: '/apiCase/apiList',
      method: 'get',
      params
    })
  }


// @Tags Case
// @Summary 分页获取Case列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Case列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiCase/caseList [get]
export const getCaseList = (params) => {
    return service({
      url: '/apiCase/caseList',
      method: 'get',
      params
    })
  }


// @Tags runCase
// @Summary 跑测试用例
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.GetById true "跑测试用例"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiCase/runApiCase [post]
export const runCase = (data) => {
  return service({
    url: '/apiCase/runApiCase',
    method: 'post',
    data
  })
}