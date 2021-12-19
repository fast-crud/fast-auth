package api

import (
	"context"
	"github.com/fast-crud/fast-auth/app/model/basic/res"
	"github.com/fast-crud/fast-auth/app/service/system"
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

type AuthController struct{}

// RegisterReq -----------------------------------------------------
type RegisterReq struct {
	g.Meta   `path:"/register" method:"post"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}
type RegisterRes struct {
	Id uint
}

func (AuthController) Register(ctx context.Context, req *RegisterReq) (res *RegisterRes, err error) {
	var info = system.UserRegisterParams{
		Avatar:   req.Avatar,
		Username: req.Username,
		Password: req.Password,
		NickName: req.NickName,
	}
	data, err := system.UserService.Register(&info)
	if err != nil {
		return nil, errors.Wrap(err, "注册失败")
	}
	return &RegisterRes{data.Id}, nil
}

// LoginReq -----------------------------------------------------
type LoginReq struct {
	g.Meta    `path:"/login" method:"post" auth:"false" per:"false"`
	Captcha   string `json:"captcha" example:"验证码"`
	Username  string `json:"username" example:"用户名"`
	Password  string `json:"password" example:"密码"`
	CaptchaId string `json:"captchaId" example:"验证码id"`
}

func (AuthController) Login(ctx context.Context, req *LoginReq) (res *res.AccessTokenRes, err error) {

	if global.Config.Captcha.Verification {
		if !system.Store.Verify(req.CaptchaId, req.Captcha, true) {
			return nil, errors.New("验证码错误")
		}
	}
	token, err := system.UserService.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return token, nil
}
