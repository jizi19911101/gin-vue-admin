package system

import (
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jizi19911101/gin-vue-admin/server/global"
	"github.com/jizi19911101/gin-vue-admin/server/model/system"
	"github.com/jizi19911101/gin-vue-admin/server/model/system/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCasbin
//@description: 更新casbin权限
//@param: authorityId string, casbinInfos []request.CasbinInfo
//@return: error

type CasbinService struct {
}

var CasbinServiceApp = new(CasbinService)

func (casbinService *CasbinService) UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	casbinService.ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range casbinInfos {
		cm := system.CasbinModel{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCasbinApi
//@description: API更新随动
//@param: oldPath string, newPath string, oldMethod string, newMethod string
//@return: error

func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.GVA_DB.Table("casbin_rule").Model(&system.CasbinModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPolicyPathByAuthorityId
//@description: 获取权限列表
//@param: authorityId string
//@return: pathMaps []request.CasbinInfo

func (casbinService *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ClearCasbin
//@description: 清除匹配的权限
//@param: v int, p ...string
//@return: bool

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Casbin
//@description: 持久化到数据库  引入自定义规则
//@return: *casbin.Enforcer

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.GVA_DB)
		modelPath := RedirectConfigFile(global.GVA_CONFIG.Casbin.ModelPath)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(modelPath, a)
		syncedEnforcer.AddFunction("ParamsMatch", casbinService.ParamsMatchFunc)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ParamsMatch
//@description: 自定义规则函数
//@param: fullNameKey1 string, key2 string
//@return: bool

func (casbinService *CasbinService) ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ParamsMatchFunc
//@description: 自定义规则函数
//@param: args ...interface{}
//@return: interface{}, error

func (casbinService *CasbinService) ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return casbinService.ParamsMatch(name1, name2), nil
}

// todo 提取为公共方法
const MaxRedirectCount = 5

var redirectCount = 0

// 一直往上级目录寻到文件
func RedirectConfigFile(path string) string {
	// 防止无线递归，找不到配置文件
	if redirectCount > MaxRedirectCount {
		return path
	}
	redirectCount++
	if _, err := os.Stat(path); err == nil {
		return path
	}
	return RedirectConfigFile("../" + path)
}
