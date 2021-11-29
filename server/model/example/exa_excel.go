package example

import "github/jizi19911101/gin-vue-admin/server/model/system"

type ExcelInfo struct {
	FileName string               `json:"fileName"` // 文件名
	InfoList []system.SysBaseMenu `json:"infoList"`
}
