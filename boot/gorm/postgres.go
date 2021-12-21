//go:build postgres
// +build postgres

package boot

import (
	"github.com/fast-crud/fast-auth/library/interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ interfaces.Gorm = (*_postgres)(nil)

var DbResolver = new(_postgres)

type _postgres struct {
	Resolver
}

// GetGormDialector 获取数据库的 gorm.Dialector

func (p *_postgres) GetGormDialector(dsn string) gorm.Dialector {
	return postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false,
	})
}

func (p *_postgres) GetConfigPath() string {
	return "config/config.postgres.yaml"
}
