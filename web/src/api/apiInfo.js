import service from '@/utils/request'

// @Tags ApiInfo
// @Summary 创建ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiInfo true "创建ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiInfo/createApiInfo [post]
export const createApiInfo = (data) => {
  data.params = JSON.stringify(data.params)
  return service({
    url: '/apiInfo/createApiInfo',
    method: 'post',
    data
  })
}

// @Tags ApiInfo
// @Summary 删除ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiInfo true "删除ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiInfo/deleteApiInfo [delete]
export const deleteApiInfo = (data) => {
  return service({
    url: '/apiInfo/deleteApiInfo',
    method: 'delete',
    data
  })
}

// @Tags ApiInfo
// @Summary 删除ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiInfo/deleteApiInfo [delete]
export const deleteApiInfoByIds = (data) => {
  return service({
    url: '/apiInfo/deleteApiInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags ApiInfo
// @Summary 更新ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiInfo true "更新ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apiInfo/updateApiInfo [put]
export const updateApiInfo = (data) => {
  return service({
    url: '/apiInfo/updateApiInfo',
    method: 'put',
    data
  })
}

// @Tags ApiInfo
// @Summary 用id查询ApiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ApiInfo true "用id查询ApiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apiInfo/findApiInfo [get]
export const findApiInfo = (params) => {
  return service({
    url: '/apiInfo/findApiInfo',
    method: 'get',
    params
  })
}

// @Tags ApiInfo
// @Summary 分页获取ApiInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ApiInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiInfo/getApiInfoList [get]
export const getApiInfoList = (params) => {
  return service({
    url: '/apiInfo/getApiInfoList',
    method: 'get',
    params
  })
}
