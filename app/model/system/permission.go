package system

import (
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Permission struct {
	global.Model
	ApplicationId  uint   `json:"applicationId" gorm:"comment:应用id"`
	OrganizationId uint   `json:"organizationId" gorm:"comment:组织id"`
	ParentId       uint   `json:"parentId" gorm:"comment:父权限id"`
	Title          string `json:"title" gorm:"not null;comment:资源标题;size:200"`
	Code           string `json:"Code" gorm:"not null;comment:权限代码;size:200"`
}

func (a *Permission) BeforeCreate(tx *gorm.DB) error {
	entity := Permission{ApplicationId: a.ApplicationId, OrganizationId: a.OrganizationId, Code: a.Code}
	if errors.Is(tx.Where(&entity).First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.Errorf(`权限代码(%s)已存在!`, a.Code)
}
func (a *Permission) TableName() string {
	return "a_permission"
}
