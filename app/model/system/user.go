package system

import (
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	global.Model
	ApplicationId  uint   `json:"applicationId" gorm:"comment:应用id"`
	OrganizationId uint   `json:"organizationId" gorm:"comment:组织id"`
	Avatar         string `json:"avatar" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Username       string `json:"userName" gorm:"comment:用户登录名"`
	Password       string `json:"-"  gorm:"comment:用户登录密码"`
	NickName       string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	PhoneCode      string `json:"phoneCode"  gorm:"comment:手机区号"`
	Phone          string `json:"phone"  gorm:"comment:手机号"`
	Email          string `json:"email"  gorm:"comment:邮箱"`
	Roles          []Role `json:"roles" gorm:"many2many:a_user_role;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	entity := User{OrganizationId: u.OrganizationId, Username: u.Username}
	if errors.Is(tx.Where(&entity).First(&entity).Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.Errorf(`用户名(%s)已存在!`, u.Username)
}

func (u *User) TableName() string {
	return "a_user"
}

// CompareHashAndPassword 密码检查 false 校验失败, true 校验成功
func (u *User) CompareHashAndPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

// EncryptedPassword 加密密码
func (u *User) EncryptedPassword() error {
	if byTes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil { // 加密密码
		return err
	} else {
		u.Password = string(byTes)
		return nil
	}
}
