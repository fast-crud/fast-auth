package system

import (
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Application struct {
	global.Model
	OrganizationId uint   `json:"organizationId" gorm:"comment:组织id"`
	Logo           string `json:"logo" gorm:"comment:Logo"`
	Code           string `json:"code" gorm:"comment:应用名称"`
	Name           string `json:"name" gorm:"comment:显示名称"`
	ClientId       string `json:"clientId" gorm:"comment:ClientId"`
	ClientSecret   string `json:"clientSecret" gorm:"comment:ClientSecret"`
	RedirectUris   string `json:"redirectUris" gorm:"comment:重定向url列表"`
}

func (app *Application) BeforeCreate(tx *gorm.DB) error {
	entity := Application{OrganizationId: app.OrganizationId, Code: app.Code}
	if errors.Is(tx.Where(&entity).First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.Errorf(`应用(%s)已存在!`, app.Code)
}
func (app *Application) TableName() string {
	return "a_application"
}

// CompareHashAndPassword 密码检查 false 校验失败, true 校验成功
func (app *Application) CompareHashAndSecret(secret string) bool {
	return app.ClientSecret == secret
}
