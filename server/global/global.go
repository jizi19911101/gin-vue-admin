package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/jizi19911101/gin-vue-admin/server/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/jizi19911101/gin-vue-admin/server/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB                  *gorm.DB
	GVA_REDIS               *redis.Client
	GVA_CONFIG              config.Server
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	GVA_Timer               = timer.NewTimerTask()
	GVA_Concurrency_Control = &singleflight.Group{}

	BlackCache local_cache.Cache
	Validate   *validator.Validate
)

func init() {
	GVA_VP = Viper()        // 初始化Viper
	GVA_LOG = Zap()         // 初始化zap日志库
	GVA_DB = Gorm().Debug() // gorm连接数据库
	Validate = validator.New()
}
