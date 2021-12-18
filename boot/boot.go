package boot

import (
	"github.com/fast-crud/fast-auth/app/service/system"
	boot "github.com/fast-crud/fast-auth/boot/gorm"
	"github.com/fast-crud/fast-auth/library/global"
)

func Initialize() {
	//configuration
	Viper.Initialize()
	// 日志
	Zap.Initialize()
	//数据库
	Gorm.Initialize(boot.DbResolver)
	if global.Db != nil {
		system.JwtBlacklist.LoadJwt()
	}
	if global.Config.System.UseMultipoint {
		Redis.Initialize()
	}
}
