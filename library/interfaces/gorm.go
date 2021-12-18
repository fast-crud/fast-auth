package interfaces

import (
	"time"

	"gorm.io/gorm"
)

type Gorm interface {
	// GetSources 获取主库的 gorm.Dialector 切片对象
	GetSources() (directories []gorm.Dialector)
	// GetReplicas 获取从库库的 gorm.Dialector 切片对象
	GetReplicas() (directories []gorm.Dialector)
	// GetResolver 通过主库与从库的链接组装 gorm.Plugin
	GetResolver() gorm.Plugin
	// GetGormDialector 获取数据库的 gorm.Dialector
	GetGormDialector(dsn string) gorm.Dialector
	// GetConfigPath 设置配置文件路径
	GetConfigPath() string
}

type GormConfig interface {
	IsEmpty() bool
	GetDsn() string
}

type GormConfigGeneral interface {
	GetMaxIdleConnes() int
	GetMaxOpenConnes() int
	GetConnMaxLifetime() time.Duration
	GetConnMaxIdleTime() time.Duration
}

type Search interface {
	Search() func(db *gorm.DB) *gorm.DB
}
