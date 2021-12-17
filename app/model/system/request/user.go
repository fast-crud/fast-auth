package request

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"gorm.io/gorm"
)

type UserRegister struct {
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}

func (r *UserRegister) Create() system.User {
	// length := len(r.AuthorityIds)
	//authorities := make([]system.Authority, 0, length)
	//for i := 0; i < length; i++ {
	//	authorities = append(authorities, system.Authority{AuthorityId: r.AuthorityIds[i]})
	//}
	return system.User{
		Avatar:   r.Avatar,
		Username: r.Username,
		Password: r.Password,
		NickName: r.NickName,
		//AuthorityId: r.AuthorityId,
		//Authorities: authorities,
	}
}

type UserLogin struct {
	Captcha   string `json:"captcha" example:"验证码"`
	Username  string `json:"username" example:"用户名"`
	Password  string `json:"password" example:"密码"`
	CaptchaId string `json:"captchaId" example:"验证码id"`
}

type UserFind struct {
	Id uint `json:"id" example:"7"`
}

func (r *UserFind) Search() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if r.Id != 0 {
			db = db.Where("id = ?", r.Id)
		}
		return db
	}
}

type UserUpdate struct {
	common.GetByID
	Avatar   string `json:"headerImg" example:"用户头像"`
	Username string `json:"userName" example:"用户登录名"`
	NickName string `json:"nickName" example:"用户昵称"`
}

func (r *UserUpdate) Update() system.User {
	return system.User{
		Avatar:   r.Avatar,
		Username: r.Username,
		NickName: r.NickName,
	}
}

type UserChangePassword struct {
	Id          uint
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

//type UserSetAuthority struct {
//	ID          uint   `json:"-"`
//	Uuid        string `json:"-"`
//	AuthorityId string `json:"authorityId"`
//}
//
//type UserSetAuthorities struct {
//	ID           uint     `json:"ID" example:"7"`
//	AuthorityIds []string `json:"authorityIds" example:"角色id切片"`
//}
