package system

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Casbin struct {
	PType      string `gorm:"column:p_type"`
	RoleId     string `gorm:"column:v0"`
	Permission string `gorm:"column:v1"`
}

func (c *Casbin) BeforeCreate(tx *gorm.DB) error {
	entity := Casbin{PType: c.PType, RoleId: c.RoleId, Permission: c.Permission}
	if errors.Is(tx.Where(&entity).First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.Errorf(`角色id(%s:%s)存在相同权限资源(%s)!`, c.PType, c.RoleId, c.Permission)
}

func (c *Casbin) TableName() string {
	return "casbin_rule"
}
