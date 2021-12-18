package api

import (
	"context"
	"github.com/fast-crud/fast-auth/app/controller/util"
	"github.com/fast-crud/fast-auth/app/model/basic/res"
	model "github.com/fast-crud/fast-auth/app/model/system"
	"github.com/fast-crud/fast-auth/app/service/system"
	"github.com/fast-crud/fast-auth/library/global"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

type UserController struct{}

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

func (UserController) Register(ctx context.Context, req *RegisterReq) (res *RegisterRes, err error) {
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

func (UserController) Login(ctx context.Context, req *LoginReq) (res *res.AccessTokenRes, err error) {

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

// MeReq -----------------------------------------------------
type MeReq struct {
	g.Meta `path:"/me" method:"post"`
}
type MeRes struct {
	model.User
}

func (UserController) Me(ctx context.Context, req *MeReq) (res *MeRes, err error) {
	var request = g.RequestFromCtx(ctx)
	var id = util.Claims.GetUserInfo(request).Id
	data, err := system.UserService.GetById(id)
	if err != nil {
		return nil, err
	}
	return &MeRes{User: *data}, nil
}

// UpdateReq -----------------------------------------------------
type UpdateReq struct {
	g.Meta   `path:"/update" method:"post"`
	Avatar   string `json:"avatar" example:"用户头像"`
	NickName string `json:"nickName" example:"用户昵称"`
}
type UpdateRes struct {
	model.User
}

func (UserController) Update(ctx context.Context, req *UpdateReq) (res *UpdateRes, err error) {
	var request = g.RequestFromCtx(ctx)
	var id = util.Claims.GetUserInfo(request).Id
	var info = system.UserUpdateParams{
		Id:       id,
		Avatar:   req.Avatar,
		NickName: req.NickName,
	}

	data, err := system.UserService.Update(&info)
	if err != nil {
		return nil, err
	}
	return &UpdateRes{*data}, nil
}

//
// ChangePasswordReq
// @Description:
//
type ChangePasswordReq struct {
	g.Meta      `path:"/changePassword" method:"post"`
	Password    string `json:"password" example:"密码"`
	NewPassword string `json:"newPassword" example:"新密码"`
}
type ChangePasswordRes struct {
}

//
// ChangePassword
// @Description:
// @receiver UserController
// @param ctx
// @param req
// @return *response.Response
//
func (UserController) ChangePassword(ctx context.Context, req *ChangePasswordReq) (res *ChangePasswordRes, err error) {

	var request = g.RequestFromCtx(ctx)
	var user = util.Claims.GetUserInfo(request)
	var id = user.Id
	var params = system.UserChangePasswordParams{
		Id:          id,
		Password:    req.Password,
		NewPassword: req.NewPassword,
	}
	if err := system.UserService.ChangePassword(&params); err != nil {
		return nil, err
	}
	return &ChangePasswordRes{}, nil
}
