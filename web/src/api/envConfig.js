import service from '@/utils/request'

// @Tags EnvConfig
// @Summary 创建EnvConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EnvConfig true "创建EnvConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /envConfig/createEnvConfig [post]
export const createEnvConfig = (data) => {
  return service({
    url: '/envConfig/createEnvConfig',
    method: 'post',
    data
  })
}

// @Tags EnvConfig
// @Summary 删除EnvConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EnvConfig true "删除EnvConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /envConfig/deleteEnvConfig [delete]
export const deleteEnvConfig = (data) => {
  return service({
    url: '/envConfig/deleteEnvConfig',
    method: 'delete',
    data
  })
}

// @Tags EnvConfig
// @Summary 删除EnvConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EnvConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /envConfig/deleteEnvConfig [delete]
export const deleteEnvConfigByIds = (data) => {
  return service({
    url: '/envConfig/deleteEnvConfigByIds',
    method: 'delete',
    data
  })
}

// @Tags EnvConfig
// @Summary 更新EnvConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EnvConfig true "更新EnvConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /envConfig/updateEnvConfig [put]
export const updateEnvConfig = (data) => {
  return service({
    url: '/envConfig/updateEnvConfig',
    method: 'put',
    data
  })
}

// @Tags EnvConfig
// @Summary 用id查询EnvConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EnvConfig true "用id查询EnvConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /envConfig/findEnvConfig [get]
export const findEnvConfig = (params) => {
  return service({
    url: '/envConfig/findEnvConfig',
    method: 'get',
    params
  })
}

// @Tags EnvConfig
// @Summary 分页获取EnvConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EnvConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /envConfig/getEnvConfigList [get]
export const getEnvConfigList = (params) => {
  return service({
    url: '/envConfig/getEnvConfigList',
    method: 'get',
    params
  })
}
