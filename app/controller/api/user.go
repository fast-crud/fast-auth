package api

import (
	"context"
	"github.com/flipped-aurora/gf-vue-admin/app/model/basic/res"
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/auth"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

type UserController struct{}

type RegisterReq struct {
	g.Meta   `path:"/register" method:"post"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}

func (UserController) Register(ctx context.Context, req *RegisterReq) (res.CommonRes, error) {
	var info = request.UserRegister{
		Avatar:   req.Avatar,
		Username: req.Username,
		Password: req.Password,
		NickName: req.NickName,
	}
	data, err := system.UserService.Register(&info)
	if err != nil {
		return res.Error("注册失败！", err), err
	}
	return res.Success(g.Map{"user": data}), nil
}

type LoginReq struct {
	g.Meta    `path:"/login" method:"post"`
	Captcha   string `json:"captcha" example:"验证码"`
	Username  string `json:"username" example:"用户名"`
	Password  string `json:"password" example:"密码"`
	CaptchaId string `json:"captchaId" example:"验证码id"`
}

func (UserController) Login(ctx context.Context, req *LoginReq) (res *res.AccessToken, err error) {

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

type MeReq struct {
	g.Meta `path:"/me" method:"post"`
}

func (UserController) Me(ctx context.Context, req *MeReq) (res.CommonRes, error) {
	var request = g.RequestFromCtx(ctx)
	var id = auth.Claims.GetUserInfo(request).Id
	data, err := system.UserService.GetById(id)
	if err != nil {
		return res.Error("获取用户信息失败!", err), err
	}
	return res.Success(g.Map{"userInfo": data}), nil
}

type UpdateReq struct {
	g.Meta   `path:"/update" method:"post"`
	Avatar   string `json:"avatar" example:"用户头像"`
	NickName string `json:"nickName" example:"用户昵称"`
}

//
// Update
// @Description:
// @receiver UserController
// @param ctx
// @param req
// @return *response.Response
//
func (UserController) Update(ctx context.Context, req *UpdateReq) (res.CommonRes, error) {
	var request = g.RequestFromCtx(ctx)
	var id = auth.Claims.GetUserInfo(request).Id
	var info = system.UserUpdateParams{
		Id:       id,
		Avatar:   req.Avatar,
		NickName: req.NickName,
	}

	data, err := system.UserService.Update(&info)
	if err != nil {
		return res.Error("更新失败", err), err
	}
	return res.Success(data), nil
}

type ChangePasswordReq struct {
	g.Meta      `path:"/changePassword" method:"post"`
	Password    string `json:"password" example:"密码"`
	NewPassword string `json:"newPassword" example:"新密码"`
}

//
// ChangePassword
// @Description:
// @receiver UserController
// @param ctx
// @param req
// @return *response.Response
//
func (UserController) ChangePassword(ctx context.Context, req *ChangePasswordReq) (res.CommonRes, error) {

	var request = g.RequestFromCtx(ctx)
	var user = auth.Claims.GetUserInfo(request)
	var id = user.Id
	var params = system.UserChangePasswordParams{
		Id:          id,
		Password:    req.Password,
		NewPassword: req.NewPassword,
	}
	if err := system.UserService.ChangePassword(&params); err != nil {
		return res.Error("修改密码失败", err), err
	}
	return res.Success(nil), nil
}
