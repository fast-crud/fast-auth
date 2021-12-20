package system

import (
	"github.com/fast-crud/fast-auth/library/global"
)

type ApplicationOauthProvider struct {
	global.Model
	ApplicationId  uint   `json:"applicationId" gorm:"comment:应用id"`
	OrganizationId uint   `json:"organizationId" gorm:"comment:组织id"`
	Provider       string `json:"provider" gorm:"comment:第三方授权类型"`
	NotBoundPolicy string `json:"notBoundPolicy" gorm:"comment:未绑定时的策略，user：由用户选择，bind：绑定已有账户，create：创建新账户"`
}

func (u *ApplicationOauthProvider) TableName() string {
	return "a_application_oauth_provider"
}
