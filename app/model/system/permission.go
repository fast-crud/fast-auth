package system

import "github.com/flipped-aurora/gf-vue-admin/library/global"

type Permission struct {
	global.Model
	ParentId   uint   `json:"parentId" gorm:"comment:父资源id"`
	Title      string `json:"title" gorm:"not null;comment:资源标题;size:200"`
	Permission string `json:"permission" gorm:"not null;comment:权限代码;size:200"`
}

func (a *Permission) TableName() string {
	return "a_permission"
}
