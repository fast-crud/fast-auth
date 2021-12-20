package system

import (
	"github.com/fast-crud/fast-auth/library/global"
)

type OauthBound struct {
	global.Model
	ApplicationId  uint   `json:"applicationId" gorm:"comment:应用id"`
	OrganizationId uint   `json:"organizationId" gorm:"comment:组织id"`
	UserId         uint   `json:"userId" gorm:"comment:用户id"`
	Provider       string `json:"provider" gorm:"comment:第三方类型"`
	OpenId         string `json:"openId" gorm:"comment:第三方id"`
}

func (u *OauthBound) TableName() string {
	return "a_oauth_bound"
}
