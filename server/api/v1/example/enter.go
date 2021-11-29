package example

import "github.com/jizi19911101/gin-vue-admin/server/service"

type ApiGroup struct {
	ExcelApi
	CustomerApi
	FileUploadAndDownloadApi
}

var excelService = service.ServiceGroupApp.ExampleServiceGroup.ExcelService
var customerService = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
var fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
