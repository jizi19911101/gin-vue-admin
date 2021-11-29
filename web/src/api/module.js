import service from '@/utils/request'

// @Tags Module
// @Summary 创建Module
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Module true "创建Module"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /module/createModule [post]
export const createModule = (data) => {
  return service({
    url: '/module/createModule',
    method: 'post',
    data
  })
}

// @Tags Module
// @Summary 删除Module
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Module true "删除Module"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /module/deleteModule [delete]
export const deleteModule = (data) => {
  return service({
    url: '/module/deleteModule',
    method: 'delete',
    data
  })
}

// @Tags Module
// @Summary 删除Module
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Module"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /module/deleteModule [delete]
export const deleteModuleByIds = (data) => {
  return service({
    url: '/module/deleteModuleByIds',
    method: 'delete',
    data
  })
}

// @Tags Module
// @Summary 更新Module
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Module true "更新Module"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /module/updateModule [put]
export const updateModule = (data) => {
  return service({
    url: '/module/updateModule',
    method: 'put',
    data
  })
}

// @Tags Module
// @Summary 用id查询Module
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Module true "用id查询Module"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /module/findModule [get]
export const findModule = (params) => {
  return service({
    url: '/module/findModule',
    method: 'get',
    params
  })
}

// @Tags Module
// @Summary 分页获取Module列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Module列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /module/getModuleList [get]
export const getModuleList = (params) => {
  return service({
    url: '/module/getModuleList',
    method: 'get',
    params
  })
}
