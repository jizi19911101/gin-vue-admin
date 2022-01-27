package global

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/songzhibin97/gkit/cache/local_cache"

	_ "github.com/jizi19911101/gin-vue-admin/server/packfile"
	"github.com/jizi19911101/gin-vue-admin/server/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	ConfigEnv  = "GVA_CONFIG"
	ConfigFile = "config.yaml"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		if config == "" {
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				config = utils.RedirectConfigFile(ConfigFile)
				//config = ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	// root 适配性
	// 根据root位置去找到对应迁移位置,保证root路径有效
	GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(GVA_CONFIG.JWT.ExpiresTime)),
	)
	return v
}
