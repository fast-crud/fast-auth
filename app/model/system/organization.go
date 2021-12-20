package system

import (
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Organization struct {
	global.Model
	Logo                 string `json:"logo" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:Logo"` // 用户头像
	Code                 string `json:"code" gorm:"comment:组织名称"`
	Name                 string `json:"name" gorm:"comment:显示名称"`
	PhoneCodeSupports    string `json:"phoneCodeSupports" gorm:"comment:手机区号支持，为空则支持所有"`
	RegisterTypeSupports bool   `json:"RegisterTypeSupports" gorm:"comment:注册类型支持"`
}

func (org *Organization) BeforeCreate(tx *gorm.DB) error {
	entity := Organization{Code: org.Code}
	if errors.Is(tx.Where(&entity).First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.Errorf(`组织(%s)已存在!`, org.Code)
}
func (org *Organization) TableName() string {
	return "a_organization"
}
