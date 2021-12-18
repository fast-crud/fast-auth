package system

import "github.com/fast-crud/fast-auth/library/global"

type Role struct {
	global.Model
	Name        string       `json:"name" gorm:"comment:角色名"` // 角色名
	Permissions []Permission `json:"permissions" gorm:"many2many:a_role_permission;"`
}

func (a *Role) TableName() string {
	return "a_role"
}
