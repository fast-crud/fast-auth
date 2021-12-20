package system

import (
	"github.com/fast-crud/fast-auth/library/global"
)

type Provider struct {
	global.Model
	Name     string `json:"name" gorm:"comment:名称"`
	Type     string `json:"type" gorm:"comment:类型"`
	Provider string `json:"provider" gorm:"comment:提供商"`
	Content  string `json:"content"  gorm:"comment:json配置"`
}

func (u *Provider) TableName() string {
	return "a_provider"
}
