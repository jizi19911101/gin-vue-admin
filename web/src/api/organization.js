import service from '@/utils/request'

// @Tags Organization
// @Summary 创建Organization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Organization true "创建Organization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organization/createOrganization [post]
export const createOrganization = (data) => {
  return service({
    url: '/organization/createOrganization',
    method: 'post',
    data
  })
}

// @Tags Organization
// @Summary 删除Organization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Organization true "删除Organization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /organization/deleteOrganization [delete]
export const deleteOrganization = (data) => {
  return service({
    url: '/organization/deleteOrganization',
    method: 'delete',
    data
  })
}

// @Tags Organization
// @Summary 删除Organization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Organization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /organization/deleteOrganization [delete]
export const deleteOrganizationByIds = (data) => {
  return service({
    url: '/organization/deleteOrganizationByIds',
    method: 'delete',
    data
  })
}

// @Tags Organization
// @Summary 更新Organization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Organization true "更新Organization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /organization/updateOrganization [put]
export const updateOrganization = (data) => {
  return service({
    url: '/organization/updateOrganization',
    method: 'put',
    data
  })
}

// @Tags Organization
// @Summary 用id查询Organization
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Organization true "用id查询Organization"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /organization/findOrganization [get]
export const findOrganization = (params) => {
  return service({
    url: '/organization/findOrganization',
    method: 'get',
    params
  })
}

// @Tags Organization
// @Summary 分页获取Organization列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Organization列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /organization/getOrganizationList [get]
export const getOrganizationList = (params) => {
  return service({
    url: '/organization/getOrganizationList',
    method: 'get',
    params
  })
}
