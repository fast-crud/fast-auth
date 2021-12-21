//go:build mysql
// +build mysql

package boot

import (
	"github.com/fast-crud/fast-auth/library/interfaces"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ interfaces.Gorm = (*_mysql)(nil)

var DbResolver = new(_mysql)

type _mysql struct {
	Resolver
}

// GetGormDialector 获取数据库的 gorm.Dialector

func (m *_mysql) GetGormDialector(dsn string) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN:                       dsn,  // DSN data source name
		DefaultStringSize:         191,  // string 类型字段的默认长度
		SkipInitializeWithVersion: true, // 根据版本自动配置
	})
}

func (m *_mysql) GetConfigPath() string {
	return "config/config.mysql.yaml"
}
