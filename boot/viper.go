package boot

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	boot "github.com/fast-crud/fast-auth/boot/gorm"
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/fsnotify/fsnotify"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
)

var Viper = new(_viper)

type _viper struct {
	err  error
	path string
}

const (
	ConfigEnv         = "GFVA_CONFIG"
	ConfigDevelopFile = "config/config.develop.yaml"
	ConfigTestingFile = "config/config.testing.yaml"
	ConfigStagingFile = "config/config.staging.yaml"
	ConfigProductFile = "config/config.product.yaml"
)

func (v *_viper) Initialize(path ...string) {
	if len(path) == 0 {
		flag.StringVar(&v.path, "c", "", "choose config file.")
		flag.Parse()
		if v.path == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				gmode.Set("unknown") // 设置gf mode
				switch gmode.Mode() {
				case gmode.DEVELOP:
					v.path = ConfigDevelopFile
					fmt.Println(`您现在的环境是 develop, 配置文件的路径为: `, v.path)
				case gmode.TESTING:
					v.path = ConfigTestingFile
					fmt.Println(`您现在的环境是 testing, 配置文件的路径为: `, v.path)
				case gmode.STAGING:
					v.path = ConfigStagingFile
					fmt.Println(`您现在的环境是 staging, 配置文件的路径为: `, v.path)
				case gmode.PRODUCT:
					v.path = ConfigProductFile
					fmt.Println(`您现在的环境是 product, 配置文件的路径为: `, v.path)
				default:
					if p := boot.DbResolver.GetConfigPath(); p != "" {
						v.path = p
						fmt.Println(`您正在使用 DbResolver.GetConfigPath() 方法传递的变量, 配置文件的路径为: `, v.path)
					}
				}
			} else {
				v.path = ConfigEnv
				fmt.Println(`您正在使用GVA_CONFIG环境变量, 配置文件的路径为: `, v.path)
			}
		} else {
			fmt.Println(`您正在使用命令行的-c参数传递的值, 配置文件的路径为: `, v.path)
		}
	} else {
		v.path = path[0]
		fmt.Println(`您正在使用func Viper()传递的值, 配置文件的路径为: `, v.path)
	}

	__viper := viper.New()
	__viper.SetConfigFile(v.path)
	if err := __viper.ReadInConfig(); err != nil {
		fmt.Println("读取文件失败!", g.Map{"path": v.path})
		log.Fatal(err)
	}
	__viper.WatchConfig()

	__viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(`配置文件已修改并更新,文件为: `, e.Name)
		if err := __viper.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := __viper.Unmarshal(&global.Config); err != nil {
		fmt.Println(`Json 序列化数据失败, err :`, err)
	}
	global.Viper = __viper
	//global.Config.AutoCode.Root, _ = filepath.Abs(".")
	global.JwtCache = local_cache.NewCache(local_cache.SetDefaultExpire(time.Second * time.Duration(global.Config.Jwt.ExpiresTime)))
}
