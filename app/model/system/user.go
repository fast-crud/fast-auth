package system

import (
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	global.Model
	Avatar   string `json:"avatar" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"` // 用户头像
	Username string `json:"userName" gorm:"comment:用户登录名"`                                              // 用户登录名
	Password string `json:"-"  gorm:"comment:用户登录密码"`                                                   // 用户登录密码
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                  // 用户昵称
	Roles    []Role `json:"roles" gorm:"many2many:a_user_role;"`
}

func (u *User) TableName() string {
	return "a_user"
}

// CompareHashAndPassword 密码检查 false 校验失败, true 校验成功
// Author [SliverHorn](https://github.com/SliverHorn)
func (u *User) CompareHashAndPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

// EncryptedPassword 加密密码
// Author [SliverHorn](https://github.com/SliverHorn)
func (u *User) EncryptedPassword() error {
	if byTes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil { // 加密密码
		return err
	} else {
		u.Password = string(byTes)
		return nil
	}
}
