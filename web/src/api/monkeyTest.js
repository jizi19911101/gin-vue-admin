import service from '@/utils/request'

// @Tags MonkeyTest
// @Summary 获取Devices列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "获取Devices列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/devices [get]
export const getDeviceList = (params) => {
    return service({
      baseURL: 'http://120.25.149.119:8082',
      url: '/api/v1/devices',
      method: 'get',
      params
    })
  }


  // @Tags StartMonkey
// @Summary 发起monkey测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "获取Devices列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/devices [get]
export const startMonkey = (data) => {
  return service({
    url: '/monkey/startMonkey',
    method: 'post',
    data
  })
}

