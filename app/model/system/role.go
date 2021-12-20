package system

import (
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Role struct {
	global.Model
	ApplicationId  uint         `json:"applicationId" gorm:"comment:应用id"`
	OrganizationId uint         `json:"organizationId" gorm:"comment:组织id"`
	Code           string       `json:"code" gorm:"comment:角色名"`  // 角色名
	Name           string       `json:"name" gorm:"comment:显示名称"` // 角色名
	Permissions    []Permission `json:"permissions" gorm:"many2many:a_role_permission;"`
}

func (a *Role) BeforeCreate(tx *gorm.DB) error {
	entity := Role{Code: a.Code, ApplicationId: a.ApplicationId, OrganizationId: a.OrganizationId}
	if errors.Is(tx.Where(&entity).First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.Errorf(`角色(%s)已存在!`, a.Code)
}
func (a *Role) TableName() string {
	return "a_role"
}
